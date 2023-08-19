package expert

import (
	"HOPE-backend/internal/entity/expert"
	"context"
	"fmt"
)

func (r *repository) GetExpertById(ctx context.Context, id uint64) (*expert.Expert, error) {
	var res expert.Expert
	err := r.db.GetContext(ctx, &res,
		r.db.Rebind(`SELECT e.id, e.name, e.photo, e.expertise, e.rating, e.price, e.is_available, e.title, 
       		e.bio, e.education, e.experience, array_agg(t.name) as topic_names
			FROM "expert".experts e, "counsel".topics t, "expert".expert_topics et
			WHERE e.id = et.expert_id AND t.id = et.topic_id AND e.id = ? GROUP BY e.id`),
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("[ExpertRepo.GetExpertById] Failed: %w", err)
	}

	return &res, nil
}
