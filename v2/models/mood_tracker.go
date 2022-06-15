package models

import (
	"errors"
	"time"

	"github.com/lib/pq"
)

type Emotion struct {
	ID          uint           `json:"id" db:"id"`
	Mood        string         `json:"mood" db:"mood"`
	Triggers    pq.StringArray `json:"triggers" db:"triggers"`
	Description string         `json:"description" db:"description"`
	TimeFrame   string         `json:"time_frame" db:"time_frame"`
	Date        int64          `json:"date" db:"date"`
	UserID      uint           `json:"user_id" db:"user_id"`
}

func getAvailableMood() []string {
	// Add new available mood in this array
	return []string{"Angry", "Sad", "Happy", "Flat", "Gorgeus"}
}

type MoodTrackerService interface {
	NewEmotion(NewEmotionRequest, uint) (*Emotion, error)
	ListEmotion(uint) ([]*Emotion, error)
	ListEmotionPerWeek(uint) ([]*Emotion, error)
	ListEmotionPerMonth(uint) ([]*Emotion, error)
}

type MoodTrackerRepository interface {
	Create(Emotion) (*Emotion, error)
	GetAllEmotionByUserID(uint) ([]*Emotion, error)
	GetAllEmotionByUserIDPerWeek(uint) ([]*Emotion, error)
	GetAllEmotionByUserIDPerMonth(uint) ([]*Emotion, error)
}

type NewEmotionRequest struct {
	Mood        string   `json:"mood"`
	Triggers    []string `json:"triggers"`
	Description string   `json:"description"`
	Time        int64    `json:"time"`
	Offset      int      `json:"offset"`
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
	loc := time.FixedZone("UTC", ths.Offset*60*60)
	time := time.UnixMilli(ths.Time).UTC().In(loc)

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
