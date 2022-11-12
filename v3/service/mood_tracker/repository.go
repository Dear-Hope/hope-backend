package mood_tracker

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/mood_tracker/filter"
)

type Repository interface {
	Store(model.Emotion) (*model.Emotion, error)
	GetAll(uint, filter.List) (model.Emotions, error)
	GetAllMood() (model.Moods, error)
}
