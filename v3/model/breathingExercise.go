package model

import "time"

type (
	BreathingExercise struct {
		ID             uint   `db:"id"`
		Title          string `db:"title"`
		Name           string `db:"name"`
		Repetition     int    `db:"repetition"`
		Description    string `db:"description"`
		Benefit        string `db:"benefit"`
		Implementation string `db:"implementation"`

		Items BreathingExerciseItems
	}

	BreathingExercises []BreathingExercise

	BreathingExerciseItem struct {
		ID         uint   `db:"id"`
		ExerciseID uint   `db:"exercise_id"`
		Name       string `db:"name"`
		Duration   int    `db:"duration"`
		Type       string `db:"type"`
	}

	BreathingExerciseItems []BreathingExerciseItem

	BreathingExerciseHistory struct {
		ID         uint      `db:"id"`
		ExerciseID uint      `db:"exercise_id"`
		UserID     uint      `db:"user_id"`
		UpdatedAt  time.Time `db:"updated_at"`
	}
)

func (ths BreathingExercise) TableWithSchemaName() string {
	return `"selfcare".breathing_exercises`
}

func (ths BreathingExerciseItem) TableWithSchemaName() string {
	return `"selfcare".breathing_exercise_items`
}

func (ths BreathingExerciseHistory) TableWithSchemaName() string {
	return `"selfcare".breathing_exercise_histories`
}

type (
	BreathingExerciseResponse struct {
		ID             uint   `json:"id"`
		Title          string `json:"title"`
		Name           string `json:"name"`
		Repetition     int    `json:"repetition"`
		Description    string `json:"howToDoDesc"`
		Benefit        string `json:"benefitDesc"`
		Implementation string `json:"implementationTime"`

		Items []BreathingExerciseItemResponse `json:"items"`
	}

	BreathingExerciseItemResponse struct {
		Name     string `json:"name"`
		Duration int    `json:"duration"`
		Type     string `json:"type"`
	}

	BreathingExerciseHistoryRequest struct {
		ExerciseID uint `json:"exerciseID"`
	}
)

func (ths BreathingExercise) ToBreathingExerciseResponse() *BreathingExerciseResponse {
	items := make([]BreathingExerciseItemResponse, len(ths.Items))
	for i, item := range ths.Items {
		items[i] = item.ToBreathingExerciseItemResponse()
	}

	return &BreathingExerciseResponse{
		ID:             ths.ID,
		Title:          ths.Title,
		Name:           ths.Name,
		Repetition:     ths.Repetition,
		Description:    ths.Description,
		Benefit:        ths.Benefit,
		Implementation: ths.Implementation,
		Items:          items,
	}
}

func (ths BreathingExerciseItem) ToBreathingExerciseItemResponse() BreathingExerciseItemResponse {
	return BreathingExerciseItemResponse{
		Name:     ths.Name,
		Duration: ths.Duration,
		Type:     ths.Type,
	}
}

func (ths BreathingExercises) ToListBreathingExercises() []BreathingExerciseResponse {
	res := make([]BreathingExerciseResponse, len(ths))
	for i, item := range ths {
		res[i] = *item.ToBreathingExerciseResponse()
	}

	return res
}
