package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths *service) GetLastPlayed(userID uint) (*model.BreathingExerciseResponse, *model.ServiceError) {
	breathingExercise, err := ths.repo.GetLastExercise(userID)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_SET_LAST_PLAYED_FAILED),
		}
	}

	return breathingExercise.ToBreathingExerciseResponse(), nil
}
