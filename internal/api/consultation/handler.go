package consultation

import (
	"HOPE-backend/internal/entity/consultation"
	"HOPE-backend/internal/entity/response"
	"context"
)

type service interface {
	UpdateStatus(ctx context.Context, id uint64, status consultation.Status) (bool, *response.ServiceError)
}

type Handler struct {
	svc service
}

func New(svc service) *Handler {
	return &Handler{svc: svc}
}
