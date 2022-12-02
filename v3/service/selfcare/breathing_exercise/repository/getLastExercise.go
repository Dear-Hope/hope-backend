package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetLastExercise(userID uint) (*model.BreathingExercise, error) {
	var breathingExercise model.BreathingExercise

	err := ths.db.Get(
		&breathingExercise,
		`WITH exercise AS (
			SELECT exercise_id AS id FROM `+model.BreathingExerciseHistory{}.TableWithSchemaName()+`
			WHERE user_id = $1 ORDER BY updated_at DESC LIMIT 1
		)
		SELECT id, title, name, repetition, description, benefit FROM `+breathingExercise.TableWithSchemaName()+` WHERE id = (SELECT id FROM exercise)
		`,
		userID,
	)
	if err != nil {
		log.Printf("get last played breathing exercise: %s", err.Error())
		return nil, err
	}

	var items model.BreathingExerciseItems
	err = ths.db.Select(&items, `SELECT name, duration, type FROM `+model.BreathingExerciseItem{}.TableWithSchemaName()+` WHERE exercise_id = $1`, breathingExercise.ID)
	if err != nil {
		log.Printf("get last played breathing exercise items: %s", err.Error())
		return nil, err
	}

	breathingExercise.Items = items

	return &breathingExercise, nil
}
