package auth

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/user"
	"HOPE-backend/pkg/helpers"
	"HOPE-backend/pkg/jwt"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (s *service) ChangePassword(ctx context.Context, req auth.ChangePasswordRequest) (*auth.TokenPairResponse, *response.ServiceError) {
	userString, err := helpers.Decrypt(strings.TrimPrefix(req.Key, "https://dearhope.id/reset?key="))
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusBadRequest,
			Msg:  "failed to decrypt reset key",
			Err:  fmt.Errorf("[AuthSvc.ChangePassword][010017] failed decrypting key: %v", err),
		}
	}

	var u user.User
	err = json.Unmarshal([]byte(userString), &u)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusBadRequest,
			Msg:  "invalid reset key given",
			Err:  fmt.Errorf("[AuthSvc.ChangePassword][010018] invalid key: %v", err),
		}
	}

	u.Password, err = helpers.EncryptPassword([]byte(req.NewPassword))
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorInternalServer,
			Err:  err,
		}
	}

	updatedUser, err := s.userRepo.UpdateUser(ctx, u)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusBadRequest,
			Msg:  constant.ErrorChangePasswordFailed,
			Err:  err,
		}
	}

	tokenPair, err := jwt.GenerateTokenPair(jwt.TokenClaim{
		Id:         updatedUser.Id,
		Role:       "placeholder",
		IsVerified: updatedUser.IsVerified,
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
