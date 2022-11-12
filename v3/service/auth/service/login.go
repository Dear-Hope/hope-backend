package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/auth/helper"
	"database/sql"
	"errors"
	"net/http"
)

func (ths *service) Login(req model.LoginRequest) (*model.TokenPairResponse, *model.ServiceError) {
	user, err := ths.repo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &model.ServiceError{
				Code: http.StatusNotFound,
				Err:  errors.New(constant.ERROR_USER_NOT_FOUND),
			}
		}
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_USER_FAILED),
		}
	}

	err = helper.ComparePassword([]byte(req.Password), []byte(user.Password))
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ERROR_PASSWORD_NOT_MATCH),
		}
	}

	if !user.IsActive {
		return nil, &model.ServiceError{
			Code: http.StatusUnauthorized,
			Err:  errors.New(constant.ERROR_ACCOUNT_NOT_ACTIVATED),
		}
	}

	tokenPair, err := helper.GenerateTokenPair(user.UserID, user.ProfileID, user.IsActive)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GENERATE_TOKEN),
		}
	}

	return tokenPair, nil
}
