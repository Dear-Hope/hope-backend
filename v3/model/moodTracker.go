package model

import (
	"HOPE-backend/v3/constant"
	"errors"
	"time"

	"github.com/lib/pq"
)

type (
	Emotion struct {
		ID          uint           `db:"id"`
		Mood        string         `db:"mood"`
		Scale       uint           `db:"scale"`
		Triggers    pq.StringArray `db:"triggers"`
		Feelings    pq.StringArray `db:"feelings"`
		Description string         `db:"description"`
		TimeFrame   string         `db:"time_frame"`
		Date        int64          `db:"date"`
		UserID      uint           `db:"user_id"`
	}

	Emotions []Emotion

	Mood string
)

func (ths Emotion) TableWithSchemaName() string {
	return `"moodtracker".emotions`
}

type (
	EmotionResponse struct {
		ID          uint     `json:"id"`
		Mood        string   `json:"mood"`
		Scale       uint     `json:"scale"`
		Triggers    []string `json:"triggers"`
		Feelings    []string `json:"feelings"`
		Description string   `json:"description"`
		TimeFrame   string   `json:"time_frame"`
		Date        int64    `json:"date"`
	}

	NewEmotionRequest struct {
		Mood        Mood     `json:"mood"`
		Scale       uint     `json:"scale"`
		Triggers    []string `json:"triggers"`
		Feelings    []string `json:"feelings"`
		Description string   `json:"description"`
		Time        int64    `json:"time"`
		Offset      int      `json:"offset"`
		UserID      uint     `json:"user_id"`
	}
)

func (ths Emotion) ToEmotionResponse() *EmotionResponse {
	return &EmotionResponse{
		ID:          ths.ID,
		Mood:        ths.Mood,
		Scale:       ths.Scale,
		Triggers:    ths.Triggers,
		Feelings:    ths.Feelings,
		Description: ths.Description,
		TimeFrame:   ths.TimeFrame,
		Date:        ths.Date,
	}
}

func (ths Emotions) ToEmotionListResponse() []EmotionResponse {
	res := make([]EmotionResponse, len(ths))
	for i, emotion := range ths {
		res[i] = *emotion.ToEmotionResponse()
	}

	return res
}

func (ths Mood) IsMoodAvailable() bool {
	// add new available mood here in the array
	availableMood := []Mood{"Angry", "Sad", "Happy", "Flat", "Gorgeus"}
	for _, mood := range availableMood {
		if ths == mood {
			return true
		}
	}

	return false
}

func (ths NewEmotionRequest) ConvertIntoTimeFrame() (string, error) {
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
		return "", errors.New(constant.ERROR_FAILED_TO_CONVERT_HOUR)
	}
}
