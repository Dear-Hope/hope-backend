package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"fmt"
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
			Err:  fmt.Errorf(constant.ERROR_SET_LAST_PLAYED_FAILED, "breathing exercise"),
		}
	}

	return nil
}
