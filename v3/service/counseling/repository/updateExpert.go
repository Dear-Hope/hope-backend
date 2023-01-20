package repository

import (
	"HOPE-backend/v3/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"strings"
)

func (ths *repository) UpdateExpert(updatedExpert model.Expert, updatedTopicIds, deletedTopicIds []uint) (*model.Expert, error) {
	var (
		insertValue      string
		deleteTopicQuery string
		args             []interface{}
	)
	for _, id := range updatedTopicIds {
		insertValue += fmt.Sprintf("(%d, %d),", updatedExpert.ID, id)
	}
	insertTopicQuery := fmt.Sprintf(`WITH updated AS (
		    INSERT INTO "counseling".expert_topics(expert_id, topic_id) VALUES 
			%s ON CONFLICT DO NOTHING
		),`, strings.TrimSuffix(insertValue, ","))

	if len(deletedTopicIds) > 0 {
		deleteTopicQuery = `deleted AS (
		    DELETE FROM "counseling".expert_topics WHERE expert_id = ? AND topic_id = ANY(?)
		),`
		args = append(args, updatedExpert.ID, pq.Array(deletedTopicIds))
	}

	query := fmt.Sprintf(`%s %s row AS (UPDATE "counseling".experts
			 SET name = :name,
				 expertise = :expertise,
				 price = :price,
				 title = :title,
				 education = :education,
				 experience = :experience,
				 photo = :photo,
				 bio = :bio,
				 updated_at = now()
			 WHERE id = ? AND is_deleted = false RETURNING id
		 ) 
		 SELECT array_agg(ts.name) AS topic_names from row r, "counseling".topics ts
		 WHERE ts.id = ANY(?)
		 GROUP BY r.id`, insertTopicQuery, deleteTopicQuery,
	)

	q, a, err := sqlx.Named(query, updatedExpert)
	if err != nil {
		log.Printf("expert update failed: %s", err.Error())
		return nil, err
	}
	a = append(a, updatedExpert.ID, pq.Array(updatedTopicIds))
	args = append(args, a...)

	rows, err := ths.db.Queryx(ths.db.Rebind(q), args...)
	if err != nil {
		log.Printf("expert update failed: %s", err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		if err = rows.Scan(&updatedExpert.Topics); err != nil {
			log.Printf("new expert create failed: %s", err.Error())
			return nil, err
		}
	}

	return &updatedExpert, nil
}
