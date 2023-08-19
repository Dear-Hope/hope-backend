package expert

import (
	"context"
	"fmt"
)

func (r *repository) GetExpertTopics(ctx context.Context, expertId uint64) ([]uint64, error) {
	var topicIds []uint64
	err := r.db.SelectContext(ctx, &topicIds,
		r.db.Rebind(`SELECT topic_id FROM "expert".expert_topics WHERE expert_id = ?`),
		expertId,
	)
	if err != nil {
		return nil, fmt.Errorf("[ExpertRepo.GetExpertTopics] Failed: %w", err)
	}

	return topicIds, nil
}
