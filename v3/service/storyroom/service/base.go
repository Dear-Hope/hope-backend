package service

import (
	"HOPE-backend/v3/service/auth"
	"HOPE-backend/v3/service/storyroom"
)

type service struct {
	repo     storyroom.Repository
	authRepo auth.Repository
}

func NewService(repo storyroom.Repository, authRepo auth.Repository) storyroom.Service {
	return &service{
		repo:     repo,
		authRepo: authRepo,
	}
}
