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
	GetTimeslotsByDay(ctx context.Context, expertId, typeId uint64, day string) (schedule.Timeslots, error)
}

type consulRepository interface {
	GetMapBookedConsultation(ctx context.Context, expertId uint64, date string) (map[string]bool, error)
}

type service struct {
	repo       repository
	consulRepo consulRepository
}

func New(repo repository, consulRepo consulRepository) *service {
	return &service{repo: repo, consulRepo: consulRepo}
}
