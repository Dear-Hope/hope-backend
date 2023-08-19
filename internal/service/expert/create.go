package expert

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/expert"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/pkg/helpers"
	"HOPE-backend/pkg/jwt"
	"context"
	"errors"
	"net/http"
	"strings"
)

func (s *service) Create(ctx context.Context, req expert.CreateUpdateRequest) (
	*auth.TokenPairResponse, *response.ServiceError) {
	if len(req.TopicIds) <= 0 {
		return nil, &response.ServiceError{
			Code: http.StatusBadRequest,
			Msg:  constant.ErrorExpertMinimalOneTopic,
			Err:  errors.New("[ExpertSvc] error no topic"),
		}
	}

	hashedPassword, err := helpers.EncryptPassword([]byte(req.Password))
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorInternalServer,
			Err:  err,
		}
	}

	req.Password = hashedPassword
	res, err := s.repo.CreateExpert(ctx, req.ToExpertModel(), req.TopicIds)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &response.ServiceError{
				Code: http.StatusBadRequest,
				Msg:  constant.ErrorExpertAlreadyExists,
				Err:  err,
			}
		} else {
			return nil, &response.ServiceError{
				Code: http.StatusInternalServerError,
				Msg:  constant.ErrorCreateExpertFailed,
				Err:  err,
			}
		}
	}

	if s.scheduleRepo.CreateSchedule(ctx, res.Id) != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorCreateExpertScheduleFailed,
			Err:  err,
		}
	}

	tokenPair, err := jwt.GenerateTokenPair(jwt.TokenClaim{
		Id:         res.Id,
		Role:       "EXPERT",
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
