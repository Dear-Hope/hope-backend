package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/auth/helper"
	"database/sql"
	"errors"
	"net/http"

	"github.com/pquerna/otp/totp"
)

func (ths *service) Activate(req model.ActivateRequest) (*model.TokenPairResponse, *model.ServiceError) {
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

	if !totp.Validate(req.Code, user.SecretKey) {
		return nil, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ERROR_OTP_CODE_EXPIRED),
		}
	}

	err = ths.repo.SetToActive(user.UserID)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_ACTIVATE_USER_FAILED),
		}
	}

	tokenPair, err := helper.GenerateTokenPair(user.UserID, user.ProfileID, true)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GENERATE_TOKEN),
		}
	}

	return tokenPair, nil
}
