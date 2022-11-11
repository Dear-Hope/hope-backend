package repository

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/mood_tracker/filter"
	"log"
)

func (ths *repository) GetAll(userID uint, f filter.List) (model.Emotions, error) {
	emotions := model.Emotions{}

	query := `SELECT e.id, m.name as mood, e.description, e.time_frame, e.date, e.triggers, e.user_id, e.feelings, e.scale
	FROM ` + model.Emotion{}.TableWithSchemaName() + ` e, ` + model.Mood{}.TableWithSchemaName() + ` m
	WHERE e.user_id = $1 AND e.mood_id = m.id `

	if f.PeriodFrom != nil {
		query += `AND e.date >= $2 `
	}

	if f.PeriodTo != nil {
		query += `AND e.date <= $3`
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
