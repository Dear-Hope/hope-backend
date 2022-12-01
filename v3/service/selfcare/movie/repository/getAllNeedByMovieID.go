package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAllNeedByMovieID(movieID uint) (model.Needs, error) {
	var needs model.Needs

	query := `SELECT id, name FROM ` + model.Need{}.TableWithSchemaName() + ` WHERE id IN (
		SELECT mnd.need_id FROM "selfcare".movie_need_details mnd WHERE mnd.movie_id = $1
	)`

	err := ths.db.Select(&needs, query, movieID)
	if err != nil {
		log.Printf("get all needs movie detail: %s", err.Error())
		return nil, err
	}

	return needs, nil
}
