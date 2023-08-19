package expert

import (
	"HOPE-backend/internal/entity/expert"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strings"
	"time"
)

func (r *repository) UpdateExpert(ctx context.Context, updatedExpert expert.Expert, updatedTopicIds,
	deletedTopicIds []uint64) (*expert.Expert, error) {
	var (
		insertValue      string
		insertTopicQuery string
		deleteTopicQuery string
		args             []interface{}
	)
	if len(updatedTopicIds) > 0 {
		for _, id := range updatedTopicIds {
			insertValue += fmt.Sprintf("(%d, %d),", updatedExpert.Id, id)
		}
		insertTopicQuery = fmt.Sprintf(`updated AS (
		    INSERT INTO "expert".expert_topics(expert_id, topic_id) VALUES 
			%s ON CONFLICT DO NOTHING
		),`, strings.TrimSuffix(insertValue, ","))

	}

	if len(deletedTopicIds) > 0 {
		deleteTopicQuery = `deleted AS (
		    DELETE FROM "expert".expert_topics WHERE expert_id = ? AND topic_id = ANY(?)
		),`
		args = append(args, updatedExpert.Id, pq.Array(deletedTopicIds))
	}

	updatedExpert.UpdatedAt = time.Now().UTC()
	query := fmt.Sprintf(`WITH %s %s row AS (UPDATE "expert".experts
			SET
				email = COALESCE(NULLIF(:email, ''), email),
		    	password = COALESCE(NULLIF(:password, ''), password),
				name = COALESCE(NULLIF(:name, ''), name),
				expertise = COALESCE(NULLIF(:expertise, ''), expertise),
				price = COALESCE(NULLIF(:price, 0), price),
				title = COALESCE(NULLIF(:title, ''), title),
				education = COALESCE(NULLIF(:education, ''), education),
				experience = COALESCE(NULLIF(:experience, ''), experience),
				photo = COALESCE(NULLIF(:photo, ''), photo),
				bio = COALESCE(NULLIF(:bio, ''), bio),
				updated_at = :updated_at
			 WHERE id = ? AND is_deleted = false RETURNING id
		 ) 
		 SELECT array_agg(ts.name) AS topic_names from row r, "counsel".topics ts
		 WHERE ts.id = ANY(?)
		 GROUP BY r.id`, insertTopicQuery, deleteTopicQuery,
	)

	q, a, err := sqlx.Named(query, updatedExpert)
	if err != nil {
		return nil, fmt.Errorf("[ExpertRepo.UpdateExpert] Failed named: %w", err)
	}
	a = append(a, updatedExpert.Id, pq.Array(updatedTopicIds))
	args = append(args, a...)

	rows, err := r.db.QueryxContext(ctx, r.db.Rebind(q), args...)
	if err != nil {
		return nil, fmt.Errorf("[ExpertRepo.UpdateExpert] Failed exec: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		if err = rows.Scan(&updatedExpert.Topics); err != nil {
			return nil, fmt.Errorf("[ExpertRepo.UpdateExpert] Failed scan: %w", err)
		}
	}

	return &updatedExpert, nil
}
