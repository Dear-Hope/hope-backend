package breathing_exercise

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
)

type Repository interface {
	GetAll(f filter.ListExercise) (model.BreathingExercises, error)
	StoreHistory(newHistory model.BreathingExerciseHistory) (*model.BreathingExerciseHistory, error)
	GetLastExercise(userID uint) (*model.BreathingExercise, error)
}
