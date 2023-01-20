package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"net/http"
)

func (ths *service) GetExpert(id uint) (*model.ExpertResponse, *model.ServiceError) {
	expert, err := ths.repo.GetExpertById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &model.ServiceError{
				Code: http.StatusNotFound,
				Err:  errors.New(constant.ErrorExpertNotFound),
			}
		}
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ErrorGetExpertFailed),
		}
	}

	return expert.ToExpertResponse(), nil
}
