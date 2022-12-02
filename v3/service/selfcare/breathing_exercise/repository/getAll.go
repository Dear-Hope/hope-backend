package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAll() (model.BreathingExercises, error) {
	var breathingExercises model.BreathingExercises

	query := `SELECT id, title, name, repetition, description, benefit FROM ` + model.BreathingExercise{}.TableWithSchemaName()

	err := ths.db.Select(&breathingExercises, query)
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
