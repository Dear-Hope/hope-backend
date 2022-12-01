package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetByID(id uint) (*model.Movie, error) {
	var movie model.Movie

	err := ths.db.Get(
		&movie,
		`SELECT id, title, year, country, genres, description, poster_link, trailer_link 
		FROM `+movie.TableWithSchemaName()+` 
		WHERE id = $1`,
		id,
	)
	if err != nil {
		log.Printf("get detail movie: %s", err.Error())
		return nil, err
	}

	return &movie, nil
}
