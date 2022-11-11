package repository

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/mood_tracker/filter"
	"log"
)

func (ths *repository) GetAll(userID uint, f filter.List) (model.Emotions, error) {
	emotions := model.Emotions{}

	query := `SELECT id, mood, description, time_frame, date, triggers, user_id, feelings, scale
	FROM "moodtracker".emotions
	WHERE user_id = $1 `

	if f.PeriodFrom != nil {
		query += `AND date >= $2 `
	}

	if f.PeriodTo != nil {
		query += `AND date <= $3`
	}

	err := ths.db.Select(
		&emotions,
		query,
		userID, *f.PeriodFrom, *f.PeriodTo,
	)
	if err != nil {
		log.Printf("emotions get all by patient id per week: %s", err.Error())
		return nil, err
	}

	return emotions, nil
}
