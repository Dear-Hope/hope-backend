package auth

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/response"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/pquerna/otp/totp"
	"net/http"
)

func (s *service) Verify(ctx context.Context, req auth.VerifyRequest) (*auth.TokenPairResponse, *response.ServiceError) {
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

	if !totp.Validate(req.Code, user.SecretKey) {
		return nil, &response.ServiceError{
			Code: http.StatusBadRequest,
			Msg:  constant.ErrorOtpCodeExpired,
			Err:  fmt.Errorf("[AuthSvc.Verify][010012] otp code has expired"),
		}
	}

	err = s.repo.VerifyUser(ctx, user.Id)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorVerifyUserFailed,
			Err:  err,
		}
	}

	tokenPair, err := generateTokenPair(user.Id, user.Role, true)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGenerateToken,
			Err:  err,
		}
	}

	return tokenPair, nil
}
