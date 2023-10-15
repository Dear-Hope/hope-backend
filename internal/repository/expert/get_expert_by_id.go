package expert

import (
	"HOPE-backend/internal/entity/expert"
	"context"
	"fmt"
)

func (r *repository) GetExpertById(ctx context.Context, id uint64) (*expert.Expert, error) {
	var res expert.Expert
	err := r.db.GetContext(ctx, &res,
		r.db.Rebind(`SELECT e.id, e.name, e.photo, e.expertise, e.price, e.is_available, e.title,
       		e.bio, e.education, e.experience, array_agg(t.name) as topic_names,
       		COALESCE(AVG(er.rating), 0) as rating FROM "expert".experts e
			INNER JOIN "expert".expert_topics et ON e.id = et.expert_id
			INNER JOIN "counsel".topics t ON t.id = et.topic_id
            LEFT JOIN "expert".expert_reviews er ON e.id = er.expert_id
			WHERE e.id = ? GROUP BY e.id`),
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("[ExpertRepo.GetExpertById] Failed: %w", err)
	}

	return &res, nil
}
