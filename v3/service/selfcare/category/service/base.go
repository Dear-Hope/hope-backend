package service

import (
	"HOPE-backend/v3/service/selfcare/category"
)

type service struct {
	repo category.Repository
}

func NewService(repo category.Repository) category.Service {
	return &service{
		repo: repo,
	}
}
