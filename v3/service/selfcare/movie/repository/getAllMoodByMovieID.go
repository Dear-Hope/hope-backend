package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAllMoodByMovieID(movieID uint) (model.Moods, error) {
	var moods model.Moods

	query := `SELECT id, name FROM ` + model.Mood{}.TableWithSchemaName() + ` WHERE id IN (
		SELECT mmd.mood_id FROM "selfcare".movie_mood_details mmd WHERE mmd.movie_id = $1
	)`

	err := ths.db.Select(&moods, query, movieID)
	if err != nil {
		log.Printf("get all moods movie detail: %s", err.Error())
		return nil, err
	}

	return moods, nil
}
