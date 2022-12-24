package self_healing_audio

import (
	"HOPE-backend/v3/model"
)

type Service interface {
	ListTheme() ([]model.SelfHealingAudioThemeListResponse, *model.ServiceError)
	GetTheme(id, userID uint) (*model.SelfHealingAudioThemeResponse, *model.ServiceError)
	GetAudio(id uint) (*model.SelfHealingAudioResponse, *model.ServiceError)
	SetLastPlayed(userID uint, req model.SelfHealingAudioHistoryRequest) *model.ServiceError
	GetLastPlayed(userID uint) (*model.SelfHealingAudioResponse, *model.ServiceError)
	ListAudioByMood(moodID, userID uint) ([]model.SelfHealingAudioListResponse, *model.ServiceError)
}
