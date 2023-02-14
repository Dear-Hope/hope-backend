package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/helpers"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/counseling/filter"
	"errors"
	"net/http"
)

func (ths *service) ListExpertSchedule(f filter.ListExpertSchedule) ([]model.ConsultationResponse, *model.ServiceError) {
	consuls, err := ths.repo.GetAllExpertSchedule(
		f.ExpertId,
		f.TypeId,
		helpers.GetStartOfDayUTCFromDateWithOffset(f.Date, f.Offset),
		helpers.GetEndOfDayUTCFromDateWithOffset(f.Date, f.Offset),
	)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ErrorGetListExpertScheduleFailed),
		}
	}

	return consuls.ToListConsulResponse(), nil
}
