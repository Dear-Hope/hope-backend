package user

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/user"
	"HOPE-backend/pkg/jwt"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/pquerna/otp/totp"
	"net/http"
)

func (s *service) Verify(ctx context.Context, req user.VerifyRequest) (
	*auth.TokenPairResponse, *response.ServiceError) {
	u, err := s.repo.GetUserByEmail(ctx, req.Email)
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

	if !totp.Validate(req.Code, u.SecretKey) {
		return nil, &response.ServiceError{
			Code: http.StatusBadRequest,
			Msg:  constant.ErrorOtpCodeExpired,
			Err:  fmt.Errorf("[AuthSvc.Verify][010012] otp code has expired"),
		}
	}

	err = s.repo.VerifyUser(ctx, u.Id)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorVerifyUserFailed,
			Err:  err,
		}
	}

	tokenPair, err := jwt.GenerateTokenPair(jwt.TokenClaim{
		Id:         u.Id,
		Role:       "USER",
		IsVerified: true,
	})
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGenerateToken,
			Err:  err,
		}
	}

	return tokenPair, nil
}
