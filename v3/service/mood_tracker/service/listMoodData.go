package service

import (
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths service) ListMoodData() ([]model.MoodResponse, *model.ServiceError) {
	moods, err := ths.moodRepo.GetAllMood()
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New("get all mood master data failed"),
		}
	}

	return moods.ToMoodListResponse(), nil
}
