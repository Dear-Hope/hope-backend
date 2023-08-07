package user

import (
	"HOPE-backend/internal/entity/user"
	"context"
)

type Repository interface {
	GetUserById(ctx context.Context, uid uint64) (*user.User, error)
}

type service struct {
	repo Repository
}

func New(repo Repository) *service {
	return &service{repo: repo}
}
