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
		MoodID      uint           `db:"mood_id"`
		UserID      uint           `db:"user_id"`
		Scale       uint           `db:"scale"`
		Triggers    pq.StringArray `db:"triggers"`
		Feelings    pq.StringArray `db:"feelings"`
		Description string         `db:"description"`
		TimeFrame   string         `db:"time_frame"`
		Date        int64          `db:"date"`
		Mood        string         `db:"mood"`
	}

	Emotions []Emotion

	Mood struct {
		ID   uint   `db:"id"`
		Name string `db:"name"`
	}

	Moods []Mood
)

func (ths Emotion) TableWithSchemaName() string {
	return `"moodtracker".emotions`
}

func (ths Mood) TableWithSchemaName() string {
	return `"moodtracker".moods`
}

type (
	MoodResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

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
		MoodID      uint     `json:"mood_id"`
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

func (ths Mood) ToMoodResponse() *MoodResponse {
	return &MoodResponse{
		ID:   ths.ID,
		Name: ths.Name,
	}
}

func (ths Moods) ToMoodListResponse() []MoodResponse {
	res := make([]MoodResponse, len(ths))
	for i, emotion := range ths {
		res[i] = *emotion.ToMoodResponse()
	}

	return res
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
