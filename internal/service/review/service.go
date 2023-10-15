package review

import (
	"HOPE-backend/internal/entity/review"
	"context"
)

type repository interface {
	GetReviewsByExpertId(ctx context.Context, expertId uint64) (review.Reviews, error)
}

type service struct {
	repo repository
}

func New(repo repository) *service {
	return &service{repo: repo}
}
