package expert

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/expert"
	"HOPE-backend/internal/entity/response"
	"context"
	"database/sql"
	"errors"
	"net/http"
)

func (s *service) Get(ctx context.Context, id uint64) (*expert.Response, *response.ServiceError) {
	ex, err := s.repo.GetExpertById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ServiceError{
				Code: http.StatusNotFound,
				Msg:  constant.ErrorExpertNotFound,
				Err:  err,
			}
		}
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetExpertFailed,
			Err:  err,
		}
	}

	return ex.ToResponse(), nil
}
