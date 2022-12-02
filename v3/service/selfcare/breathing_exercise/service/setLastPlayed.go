package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths *service) SetLastPlayed(userID uint, req model.BreathingExerciseHistoryRequest) *model.ServiceError {
	_, err := ths.repo.StoreHistory(model.BreathingExerciseHistory{
		ExerciseID: req.ExerciseID,
		UserID:     userID,
	})
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_SET_LAST_PLAYED_FAILED),
		}
	}

	return nil
}
