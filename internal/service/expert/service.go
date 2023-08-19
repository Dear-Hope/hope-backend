package expert

import (
	"HOPE-backend/internal/entity/expert"
	"context"
)

type repository interface {
	CreateExpert(ctx context.Context, expert expert.Expert, topicsId []uint64) (*expert.Expert, error)
	GetExpertByEmail(ctx context.Context, email string) (*expert.Expert, error)
	GetExpertById(ctx context.Context, id uint64) (*expert.Expert, error)
	UpdateExpert(ctx context.Context, updatedExpert expert.Expert, updatedTopicIds,
		deletedTopicIds []uint64) (*expert.Expert, error)
	GetExpertTopics(ctx context.Context, expertId uint64) ([]uint64, error)
}

type scheduleRepository interface {
	CreateSchedule(ctx context.Context, expertId uint64) error
}

type service struct {
	repo         repository
	scheduleRepo scheduleRepository
}

func New(repo repository, schRepo scheduleRepository) *service {
	return &service{repo: repo, scheduleRepo: schRepo}
}
