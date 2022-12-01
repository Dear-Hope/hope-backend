package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) Store(newMovie model.Movie) (*model.Movie, error) {
	rows, err := ths.db.NamedQuery(
		`INSERT INTO `+newMovie.TableWithSchemaName()+` (title, year, country, genres, description, poster_link, trailer_link)
		VALUES (:title, :year, :country, :genres, :description, :poster_link, :trailer_link)
		RETURNING id`,
		newMovie,
	)
	if err != nil {
		log.Printf("new movie create failed: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&newMovie.ID); err != nil {
			log.Printf("new movie create failed: %s", err.Error())
			return nil, err
		}
	}

	return &newMovie, nil
}
