package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths *service) BookExpertSchedule(req model.BookingRequest) (bool, *model.ServiceError) {
	success, err := ths.repo.UpdateExpertScheduleUser(req.Id, req.UserId)
	if err != nil {
		return false, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ErrorBookExpertScheduleFailed),
		}
	}

	return success, nil
}
