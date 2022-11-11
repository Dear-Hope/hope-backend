package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/auth/helper"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (ths *service) ChangePassword(req model.ChangePasswordRequest) (*model.TokenPairResponse, *model.ServiceError) {
	userString, err := helper.Decrypt(strings.TrimPrefix(req.Key, "https://dearhope.id/reset?key="))
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New("failed to decrypt change password key: " + err.Error()),
		}
	}

	var user model.User
	err = json.Unmarshal([]byte(userString), &user)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New("invalid key: " + err.Error()),
		}
	}

	hashedPassword, err := helper.EncryptPassword([]byte(req.NewPassword))
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ERROR_CHANGE_PASSWORD_FAILED),
		}
	}

	err = ths.repo.UpdatePassword(user.UserID, hashedPassword)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ERROR_CHANGE_PASSWORD_FAILED),
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
