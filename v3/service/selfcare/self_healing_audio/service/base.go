package service

import (
	"HOPE-backend/v3/service/selfcare/self_healing_audio"
)

type service struct {
	repo self_healing_audio.Repository
}

func NewService(repo self_healing_audio.Repository) self_healing_audio.Service {
	return &service{
		repo: repo,
	}
}
