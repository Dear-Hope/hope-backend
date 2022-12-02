package service

import (
	"HOPE-backend/v3/service/selfcare/breathing_exercise"
)

type service struct {
	repo breathing_exercise.Repository
}

func NewService(repo breathing_exercise.Repository) breathing_exercise.Service {
	return &service{
		repo: repo,
	}
}
