package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/counseling/filter"
	"errors"
	"net/http"
)

func (ths *service) ListExperts(f filter.ListExpert) ([]model.ExpertResponse, *model.ServiceError) {
	experts, err := ths.repo.GetAllExperts(f)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ErrorGetListExpertsFailed),
		}
	}

	return experts.ToListExpertResponse(), nil
}
