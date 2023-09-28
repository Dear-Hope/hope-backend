package consultation

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/consultation"
	"HOPE-backend/internal/entity/response"
	"context"
	"net/http"
)

func (s *service) GetDetailByExpert(ctx context.Context, consulId uint64) (*consultation.ExpertResponse,
	*response.ServiceError) {

	consul, err := s.repo.GetConsultationById(ctx, consulId)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetDetailConsultationFailed,
			Err:  err,
		}
	}

	return s.constructExpertResponse(ctx, *consul)
}
