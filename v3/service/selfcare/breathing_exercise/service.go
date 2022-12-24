package breathing_exercise

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
)

type Service interface {
	List(f filter.ListExercise) ([]model.BreathingExerciseResponse, *model.ServiceError)
	SetLastPlayed(userID uint, req model.BreathingExerciseHistoryRequest) *model.ServiceError
	GetLastPlayed(userID uint) (*model.BreathingExerciseResponse, *model.ServiceError)
}
