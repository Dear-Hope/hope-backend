package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"net/http"
)

func (ths *service) GetUpcomingSchedule(expertId, typeId int64) (*model.ConsultationResponse, *model.ServiceError) {
	consul, err := ths.repo.GetExpertUpcomingSchedule(expertId, typeId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &model.ServiceError{
				Code: http.StatusNotFound,
				Err:  errors.New(constant.ErrorUpcomingScheduleNotFound),
			}
		}
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ErrorGetExpertUpcomingScheduleFailed),
		}
	}

	return consul.ToConsulResponse(), nil
}
