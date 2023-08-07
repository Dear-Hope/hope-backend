package user

import (
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/user"
	"context"
)

type service interface {
	GetUser(ctx context.Context, userId uint64) (*user.Response, *response.ServiceError)
}

type Handler struct {
	svc service
}

func New(svc service) *Handler {
	return &Handler{svc: svc}
}
