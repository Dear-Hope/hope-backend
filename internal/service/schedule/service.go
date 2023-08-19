package schedule

import (
	"HOPE-backend/internal/entity/schedule"
	"context"
)

type repository interface {
	GetSchedulesByExpertId(ctx context.Context, expertId uint64) (schedule.Schedules, error)
	GetTimeslotsByScheduleIds(ctx context.Context, timeslotIds []uint64) (schedule.Timeslots, error)
	UpdateSchedule(ctx context.Context, req schedule.Schedule) (*schedule.Schedule, error)
	UpdateTimeslots(ctx context.Context, req schedule.Timeslots) (*schedule.Timeslots, error)
	DeleteTimeslots(ctx context.Context, deletedIds []uint64) error
}

type service struct {
	repo repository
}

func New(repo repository) *service {
	return &service{repo: repo}
}
