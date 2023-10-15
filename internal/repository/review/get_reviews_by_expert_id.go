package review

import (
	"HOPE-backend/internal/entity/review"
	"context"
	"fmt"
)

func (r *repository) GetReviewsByExpertId(ctx context.Context, expertId uint64) (review.Reviews, error) {
	var reviews review.Reviews
	err := r.db.SelectContext(ctx, &reviews,
		r.db.Rebind(`SELECT id, expert_id, rating, review, created_at, updated_at, is_deleted 
			FROM "expert".expert_reviews WHERE is_deleted = false AND expert_id = ?`),
		expertId,
	)
	if err != nil {
		return nil, fmt.Errorf("[ReviewRepo.GetReviewsByExpertId] Failed: %w", err)
	}

	return reviews, nil
}
