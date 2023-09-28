package consultation

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/consultation"
	"HOPE-backend/internal/entity/response"
	"context"
	"errors"
	"net/http"
)

func (s *service) UpdateStatus(ctx context.Context, id uint64, status consultation.Status) (bool, *response.ServiceError) {
	if status <= 0 {
		return false, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorInvalidStatusConsultation,
			Err:  errors.New("[ConsultationSvc.UpdateStatus] invalid status given"),
		}
	}

	success, err := s.repo.UpdateStatusConsultation(ctx, id, status.String())
	if err != nil {
		return false, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorUpdateStatusConsultationFailed,
			Err:  err,
		}
	}

	return success, nil
}
