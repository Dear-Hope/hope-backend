package user

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/user"
	"context"
	"database/sql"
	"errors"
	"net/http"
)

func (s *service) GetUser(ctx context.Context, uid uint64) (*user.Response, *response.ServiceError) {
	result, err := s.repo.GetUserById(ctx, uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ServiceError{
				Code: http.StatusNotFound,
				Msg:  constant.ErrorUserNotFound,
				Err:  err,
			}
		}
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetUserFailed,
			Err:  err,
		}
	}

	return result.ToResponse(), nil
}
