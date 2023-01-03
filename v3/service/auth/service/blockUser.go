package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths *service) BlockUser(userID uint, req model.BlockUserRequest) *model.ServiceError {
	err := ths.repo.CreateBlockedUser(userID, req.BlockedUserID)
	if err != nil {
		return &model.ServiceError{
			Err:  errors.New(constant.ERROR_BLOCK_USER_FAILED),
			Code: http.StatusInternalServerError,
		}
	}

	return nil
}
