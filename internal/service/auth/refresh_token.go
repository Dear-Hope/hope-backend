package auth

import (
	"HOPE-backend/config"
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/pkg/jwt"
	"context"
	"fmt"
	"net/http"
)

func (s *service) RefreshToken(ctx context.Context, token string) (*auth.TokenPairResponse, *response.ServiceError) {
	claim, err := jwt.AuthorizeToken(token, config.Get().Server.SecretKey, true)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusUnauthorized,
			Msg:  err.Error(),
			Err:  fmt.Errorf("[AuthSvc.RefreshToken][010020] failed to authorize token: %v", err),
		}
	}

	tokenPair, err := jwt.GenerateTokenPair(jwt.TokenClaim{
		Id:         claim.Id,
		Role:       claim.Role,
		IsVerified: claim.IsVerified,
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
