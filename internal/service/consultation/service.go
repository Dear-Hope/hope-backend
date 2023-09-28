package consultation

import (
	"HOPE-backend/internal/entity/consultation"
	"HOPE-backend/internal/entity/user"
	"context"
)

type repository interface {
	GetAllConsultation(ctx context.Context, req consultation.Consultation, isFilterPerMonth bool) (
		consultation.Consultations, error)
	GetConsultationById(ctx context.Context, consulId uint64) (*consultation.Consultation, error)
	UpdateStatusConsultation(ctx context.Context, id uint64, status string) (bool, error)
}

type userRepository interface {
	GetUserById(ctx context.Context, uid uint64) (*user.User, error)
}

type service struct {
	repo     repository
	userRepo userRepository
}

func New(repo repository, userRepo userRepository) *service {
	return &service{repo: repo, userRepo: userRepo}
}
