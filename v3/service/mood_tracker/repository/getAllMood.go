package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAllMood() (model.Moods, error) {
	moods := model.Moods{}

	err := ths.db.Select(
		&moods,
		`SELECT id, name FROM "moodtracker".moods`,
	)
	if err != nil {
		log.Printf("get all mood master data: %s", err.Error())
		return nil, err
	}

	return moods, nil
}
