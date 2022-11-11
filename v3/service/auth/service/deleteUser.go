package service

import (
	"HOPE-backend/v3/model"
	"net/http"
)

func (ths *service) DeleteUser(req model.ResetPasswordRequest) *model.ServiceError {
	err := ths.repo.DeleteByEmail(req.Email)
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return nil
}
