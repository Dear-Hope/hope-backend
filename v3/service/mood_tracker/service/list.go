package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/mood_tracker/filter"
	"errors"
	"net/http"
)

func (ths *service) List(userID uint, f filter.List) ([]model.EmotionResponse, *model.ServiceError) {
	emotions, err := ths.moodRepo.GetAll(userID, f)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_ALL_EMOTION_FAILED),
		}
	}

	return emotions.ToEmotionListResponse(), nil
}
