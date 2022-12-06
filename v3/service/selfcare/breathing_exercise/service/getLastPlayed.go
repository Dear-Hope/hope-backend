package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

func (ths *service) GetLastPlayed(userID uint) (*model.BreathingExerciseResponse, *model.ServiceError) {
	breathingExercise, err := ths.repo.GetLastExercise(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf(constant.ERROR_GET_LAST_PLAYED_FAILED, "breathing exercise"),
		}
	}

	return breathingExercise.ToBreathingExerciseResponse(), nil
}
