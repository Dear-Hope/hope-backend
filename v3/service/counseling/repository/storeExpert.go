package repository

import (
	"HOPE-backend/v3/model"
	"fmt"
	"log"
	"strings"
)

func (ths *repository) StoreExpert(newExpert model.Expert, topicIds []uint) (*model.Expert, error) {
	var insertValue string
	for _, id := range topicIds {
		insertValue += fmt.Sprintf("((SELECT id FROM expert LIMIT 1), %d),", id)
	}
	insertTopicQuery := fmt.Sprintf(`topics AS (
		    INSERT INTO "counseling".expert_topics(expert_id, topic_id) VALUES 
			%s RETURNING topic_id
		)`, strings.TrimSuffix(insertValue, ","))

	query := fmt.Sprintf(
		`WITH expert AS (
    		INSERT INTO "counseling".experts(name, expertise, price, title, education, experience, photo, bio) 
           	VALUES (:name, :expertise, :price, :title, :education,:experience, :photo, :bio)
           	RETURNING id
    	), %s
    	SELECT e.id, array_agg(ts.name) AS topic_names from expert e, topics t, "counseling".topics ts 
		WHERE t.topic_id = ts.id GROUP BY e.id`, insertTopicQuery,
	)

	rows, err := ths.db.NamedQuery(query, newExpert)
	if err != nil {
		log.Printf("new expert create failed: %s", err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		if err = rows.Scan(&newExpert.ID, &newExpert.Topics); err != nil {
			log.Printf("new expert create failed: %s", err.Error())
			return nil, err
		}
	}

	return &newExpert, nil
}
