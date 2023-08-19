package auth

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/pkg/jwt"
	"context"
	"database/sql"
	"errors"
	"net/http"
)

func (s *service) Login(ctx context.Context, req auth.LoginRequest) (*auth.TokenPairResponse, *response.ServiceError) {
	var (
		claim    jwt.TokenClaim
		password string
	)

	if req.Source == "USER" {
		user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
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

		if !user.IsVerified {
			return nil, &response.ServiceError{
				Code: http.StatusUnauthorized,
				Msg:  constant.ErrorAccountNotVerified,
				Err:  errors.New("[AuthSvc.Login][010010] account not verified yet"),
			}
		}

		password = user.Password
		claim = jwt.TokenClaim{
			Id:         user.Id,
			Role:       "USER",
			IsVerified: user.IsVerified,
		}
	} else {
		expert, err := s.expertRepo.GetExpertByEmail(ctx, req.Email)
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

		password = expert.Password
		claim = jwt.TokenClaim{
			Id:         expert.Id,
			Role:       "EXPERT",
			IsVerified: true,
		}
	}

	err := comparePassword([]byte(req.Password), []byte(password))
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusBadRequest,
			Msg:  constant.ErrorPasswordNotMatch,
			Err:  err,
		}
	}

	tokenPair, err := jwt.GenerateTokenPair(claim)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGenerateToken,
			Err:  err,
		}
	}

	return tokenPair, nil
}
