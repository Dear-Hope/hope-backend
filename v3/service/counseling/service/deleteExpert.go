package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths *service) DeleteExpert(id uint) *model.ServiceError {
	err := ths.repo.DeleteExpert(id)
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ErrorDeleteExpertFailed),
		}
	}

	return nil
}
