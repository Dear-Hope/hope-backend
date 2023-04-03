package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"net/http"
)

func (ths *service) GetExpertAvailability(id uint) (bool, *model.ServiceError) {
	isAvailable, err := ths.repo.IsExpertAvailable(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return isAvailable, &model.ServiceError{
				Code: http.StatusNotFound,
				Err:  errors.New(constant.ErrorExpertNotFound),
			}
		}
		return isAvailable, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ErrorGetExpertFailed),
		}
	}

	return isAvailable, nil
}
