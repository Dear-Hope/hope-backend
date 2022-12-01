package service

import "HOPE-backend/v3/service/selfcare/movie"

type service struct {
	repo movie.Repository
}

func NewService(repo movie.Repository) movie.Service {
	return &service{
		repo: repo,
	}
}
