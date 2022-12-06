package self_healing_audio

import (
	"HOPE-backend/v3/model"
)

type Repository interface {
	GetAllTheme() (model.SelfHealingAudioThemes, error)
	GetThemeByID(id, userID uint) (*model.SelfHealingAudioTheme, error)
	GetAudioByID(id uint) (*model.SelfHealingAudio, error)
	StoreHistory(newHistory model.SelfHealingAudioHistory) (*model.SelfHealingAudioHistory, error)
	GetLastAudio(userID uint) (*model.SelfHealingAudio, error)
}
