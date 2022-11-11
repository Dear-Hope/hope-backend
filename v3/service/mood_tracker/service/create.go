package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
	"strings"
)

func (ths *service) Create(req model.NewEmotionRequest) (*model.EmotionResponse, *model.ServiceError) {
	if !req.Mood.IsMoodAvailable() {
		return nil, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ERROR_MOOD_NOT_LISTED),
		}
	}

	timeFrame, err := req.ConvertIntoTimeFrame()
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	newEmotion := model.Emotion{
		Mood:        string(req.Mood),
		Scale:       req.Scale,
		Triggers:    req.Triggers,
		Feelings:    req.Feelings,
		Description: req.Description,
		TimeFrame:   timeFrame,
		Date:        req.Time,
		UserID:      req.UserID,
	}

	emotion, err := ths.moodRepo.Store(newEmotion)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &model.ServiceError{
				Code: http.StatusBadRequest,
				Err:  errors.New(constant.ERROR_EMOTION_ALREADY_EXISTS),
			}
		} else {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  errors.New(constant.ERROR_CREATE_EMOTION_FAILED),
			}
		}
	}

	return emotion.ToEmotionResponse(), nil
}
