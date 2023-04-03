package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
	"strings"
)

func (ths *service) CreateExpertSchedule(req model.NewConsultationRequest, expertId int64) (int64, *model.ServiceError) {
	if len(req.Schedules) <= 0 {
		return 0, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ErrorExpertMinimalOneSchedule),
		}
	}

	rowsAffected, err := ths.repo.StoreExpertSchedule(req.Schedules, expertId)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return 0, &model.ServiceError{
				Code: http.StatusBadRequest,
				Err:  errors.New(constant.ErrorExpertScheduleAlreadyExists),
			}
		} else {
			return 0, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  errors.New(constant.ErrorCreateExpertScheduleFailed),
			}
		}
	}

	return rowsAffected, nil
}
