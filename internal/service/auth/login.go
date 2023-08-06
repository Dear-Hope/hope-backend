package auth

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/response"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

func (s *service) Login(ctx context.Context, req auth.LoginRequest) (*auth.TokenPairResponse, *response.ServiceError) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
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

	err = comparePassword([]byte(req.Password), []byte(user.Password))
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusBadRequest,
			Msg:  constant.ErrorPasswordNotMatch,
			Err:  err,
		}
	}

	if !user.IsVerified {
		return nil, &response.ServiceError{
			Code: http.StatusUnauthorized,
			Msg:  constant.ErrorAccountNotActivated,
			Err:  fmt.Errorf("[AuthSvc.Login][010010] account not verified yet"),
		}
	}

	tokenPair, err := generateTokenPair(user.Id, user.Role, user.IsVerified)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGenerateToken,
			Err:  err,
		}
	}

	return tokenPair, nil
}
