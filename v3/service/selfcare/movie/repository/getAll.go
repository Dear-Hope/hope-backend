package repository

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
	"fmt"
	"log"
)

func (ths *repository) GetAll(f filter.ListMovie) (model.Movies, error) {
	var movies model.Movies

	query := `SELECT id, title, year, country, genres, description, poster_link, trailer_link FROM ` + model.Movie{}.TableWithSchemaName() + ` m WHERE id IN (
		SELECT DISTINCT mmd.movie_id 
		FROM "selfcare".movie_mood_details mmd, "selfcare".movie_need_details mnd 
		WHERE mmd.movie_id = mnd.movie_id`

	if f.MoodID != nil && *f.MoodID > 0 {
		query += fmt.Sprintf(" AND mmd.mood_id = %d", *f.MoodID)
	}

	if f.NeedID != nil && *f.NeedID > 0 {
		query += fmt.Sprintf(" AND mnd.need_id = %d", *f.NeedID)
	}

	err := ths.db.Select(
		&movies,
		query+")",
	)
	if err != nil {
		log.Printf("get all movies: %s", err.Error())
		return nil, err
	}

	return movies, nil
}
