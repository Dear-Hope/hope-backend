package expert

import (
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/expert"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/schedule"
	"context"
)

type service interface {
	Create(ctx context.Context, req expert.CreateUpdateRequest) (*auth.TokenPairResponse, *response.ServiceError)
	Get(ctx context.Context, id uint64) (*expert.Response, *response.ServiceError)
	Update(ctx context.Context, req expert.CreateUpdateRequest) (bool, *response.ServiceError)
	//Verify(ctx context.Context, req user.VerifyRequest) (*auth.TokenPairResponse, *response.ServiceError)
	//SaveProfilePhoto(ctx context.Context, req user.SaveProfilePhotoRequest) (string, *response.ServiceError)
}

type scheduleService interface {
	Get(ctx context.Context, expertId uint64) ([]schedule.Response, *response.ServiceError)
	Update(ctx context.Context, req schedule.UpdateRequest) (bool, *response.ServiceError)
}

type Handler struct {
	svc         service
	scheduleSvc scheduleService
}

func New(svc service, schSvc scheduleService) *Handler {
	return &Handler{svc: svc, scheduleSvc: schSvc}
}
