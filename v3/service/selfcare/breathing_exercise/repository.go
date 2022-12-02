package breathing_exercise

import (
	"HOPE-backend/v3/model"
)

type Repository interface {
	GetAll() (model.BreathingExercises, error)
	StoreHistory(newHistory model.BreathingExerciseHistory) (*model.BreathingExerciseHistory, error)
	GetLastExercise(userID uint) (*model.BreathingExercise, error)
}
