package mood_tracker

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/mood_tracker/filter"
)

type Service interface {
	Create(model.NewEmotionRequest) (*model.EmotionResponse, *model.ServiceError)
	List(uint, filter.List) ([]model.EmotionResponse, *model.ServiceError)
}
