package service

import "HOPE-backend/v3/service/selfcare/music"

type service struct {
	repo music.Repository
}

func NewService(repo music.Repository) music.Service {
	return &service{
		repo: repo,
	}
}
