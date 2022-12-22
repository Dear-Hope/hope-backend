package service

import (
	"HOPE-backend/v3/service/storyroom"
)

type service struct {
	repo storyroom.Repository
}

func NewService(repo storyroom.Repository) storyroom.Service {
	return &service{
		repo: repo,
	}
}
