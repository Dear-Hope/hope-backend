package breathing_exercise

import (
	"HOPE-backend/v3/model"
)

type Service interface {
	List() ([]model.BreathingExerciseResponse, *model.ServiceError)
	SetLastPlayed(userID uint, req model.BreathingExerciseHistoryRequest) *model.ServiceError
	GetLastPlayed(userID uint) (*model.BreathingExerciseResponse, *model.ServiceError)
}
