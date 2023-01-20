package service

import (
	"HOPE-backend/v3/service/counseling"
)

type service struct {
	repo counseling.Repository
}

func NewService(repo counseling.Repository) counseling.Service {
	return &service{
		repo: repo,
	}
}
