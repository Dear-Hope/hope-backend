package repository

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
	"log"
)

func (ths *repository) GetAll(f filter.ListExercise) (model.BreathingExercises, error) {
	var breathingExercises model.BreathingExercises

	query := `SELECT id, title, name, repetition, description, benefit, implementation, sub_title FROM ` + model.BreathingExercise{}.TableWithSchemaName()
	args := []interface{}{}

	if f.MoodID > 0 {
		query += " WHERE mood_id = $1"
		args = append(args, f.MoodID)
	}

	err := ths.db.Select(&breathingExercises, query, args...)
	if err != nil {
		log.Printf("get all breathing exercises: %s", err.Error())
		return nil, err
	}

	for i, breathingExercise := range breathingExercises {
		var items model.BreathingExerciseItems
		err := ths.db.Select(&items, `SELECT name, duration, type FROM `+model.BreathingExerciseItem{}.TableWithSchemaName()+` WHERE exercise_id = $1`, breathingExercise.ID)
		if err != nil {
			log.Printf("get all breathing exercise items: %s", err.Error())
			return nil, err
		}

		breathingExercises[i].Items = items
	}

	return breathingExercises, nil
}
