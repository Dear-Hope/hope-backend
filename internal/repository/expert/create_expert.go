package expert

import (
	"HOPE-backend/internal/entity/expert"
	"context"
	"fmt"
	"strings"
)

func (r *repository) CreateExpert(ctx context.Context, newExpert expert.Expert, topicIds []uint64) (
	*expert.Expert, error) {
	var insertValue string
	for _, id := range topicIds {
		insertValue += fmt.Sprintf("((SELECT id FROM expert LIMIT 1), %d),", id)
	}
	insertTopicQuery := fmt.Sprintf(`topics AS (
		    INSERT INTO "expert".expert_topics(expert_id, topic_id) VALUES 
			%s RETURNING topic_id
		)`, strings.TrimSuffix(insertValue, ","))

	query := fmt.Sprintf(
		`WITH expert AS (
    		INSERT INTO "expert".experts(email, password, name, expertise, price, 
				title, education, experience, photo, bio) 
           	VALUES (:email, :password, :name, :expertise, :price, :title, :education,:experience, :photo, :bio)
           	RETURNING id
    	), %s
    	SELECT e.id, array_agg(ts.name) AS topic_names from expert e, topics t, "counsel".topics ts 
		WHERE t.topic_id = ts.id GROUP BY e.id`, insertTopicQuery,
	)

	rows, err := r.db.NamedQueryContext(ctx, query, newExpert)
	if err != nil {
		return nil, fmt.Errorf("[ExpertRepo.CreateExpert] Failed exec: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		if err = rows.Scan(&newExpert.Id, &newExpert.Topics); err != nil {
			return nil, fmt.Errorf("[ExpertRepo.CreateExpert] Failed scan: %w", err)
		}
	}

	return &newExpert, nil
}
