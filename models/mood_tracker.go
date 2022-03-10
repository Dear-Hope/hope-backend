package models

import (
	"errors"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Emotion struct {
	gorm.Model
	Mood        string         `json:"mood" gorm:"not null"`
	Triggers    pq.StringArray `json:"triggers" gorm:"not null;type:text[]"`
	Description string         `json:"description" gorm:"not null"`
	TimeFrame   string         `json:"time_frame" gorm:"not null"`
	PatientID   uint           `json:"patient_id" gorm:"not null"`
	Patient     User           `json:"-" gorm:"constraint:OnUpdate:CASCADE;"`
}

func (Emotion) TableName() string {
	return "emotion"
}

func getAvailableMood() []string {
	// Add new available mood in this array
	return []string{"Angry", "Sad", "Happy", "Flat", "Gorgeus"}
}

type MoodTrackerService interface {
	NewEmotion(NewEmotionRequest, uint) (*Emotion, error)
	ListEmotion(uint) ([]*Emotion, error)
	ListEmotionPerWeek(uint) ([]*Emotion, error)
}

type MoodTrackerRepository interface {
	Create(Emotion) (*Emotion, error)
	GetAllEmotionByPatientID(uint) ([]*Emotion, error)
	GetAllEmotionByPatientIDPerWeek(uint) ([]*Emotion, error)
}

type NewEmotionRequest struct {
	Mood        string   `json:"mood"`
	Triggers    []string `json:"triggers"`
	Description string   `json:"description"`
	Time        int64    `json:"time"`
}

func (ths *NewEmotionRequest) IsMoodAvailable() bool {
	for _, mood := range getAvailableMood() {
		if ths.Mood == mood {
			return true
		}
	}

	return false
}

func (ths *NewEmotionRequest) ConvertIntoTimeFrame() (string, error) {
	time := time.UnixMilli(ths.Time)

	switch hour := time.Hour(); {
	case hour >= 3 && hour <= 10:
		return "Morning", nil
	case hour >= 11 && hour <= 18:
		return "Noon", nil
	case hour >= 19 || hour <= 2:
		return "Evening", nil
	default:
		return "", errors.New("failed to convert hour")
	}
}
