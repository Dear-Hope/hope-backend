package user

import (
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/user"
	"context"
)

type service interface {
	Register(ctx context.Context, req user.RegisterRequest) (*auth.TokenPairResponse, *response.ServiceError)
	Get(ctx context.Context, userId uint64) (*user.Response, *response.ServiceError)
	Update(ctx context.Context, req user.UpdateRequest) (bool, *response.ServiceError)
	Verify(ctx context.Context, req user.VerifyRequest) (*auth.TokenPairResponse, *response.ServiceError)
	SaveProfilePhoto(ctx context.Context, req user.SaveProfilePhotoRequest) (string, *response.ServiceError)
}

type Handler struct {
	svc service
}

func New(svc service) *Handler {
	return &Handler{svc: svc}
}
