package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/auth/helper"
	"errors"
	"net/http"
)

func (ths *service) UpdateLoggedInUser(req model.UpdateRequest) (*model.UserResponse, *model.ServiceError) {
	var (
		newPassword string
		err         error
	)
	if req.Password != "" {
		newPassword, err = helper.EncryptPassword([]byte(req.Password))
		if err != nil {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  errors.New(constant.ERROR_UPDATE_USER_FAILED),
			}
		}
	}

	newUser := model.User{
		UserID:   req.UserID,
		Email:    req.Email,
		Password: newPassword,
		Name:     req.Name,
		Alias:    req.Alias,
		Profile: model.Profile{
			ProfileID:    req.ProfileID,
			ProfilePhoto: req.ProfilePhoto,
			Job:          req.Profile.Job,
			Activities:   req.Profile.Activities,
		},
	}

	updatedUser, err := ths.repo.Update(newUser)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_UPDATE_USER_FAILED),
		}
	}

	return updatedUser.ToUserResponse(), nil
}
