package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths *service) List() ([]model.BreathingExerciseResponse, *model.ServiceError) {
	breathingExercises, err := ths.repo.GetAll()
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_LIST_BREATHING_EXERCISES_FAILED),
		}
	}

	return breathingExercises.ToListBreathingExercises(), nil
}
