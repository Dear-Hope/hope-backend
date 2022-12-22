package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAllCategories() (model.PostCategories, error) {
	var categories model.PostCategories

	err := ths.db.Select(
		&categories,
		`SELECT id, name FROM "storyroom".categories`,
	)
	if err != nil {
		log.Printf("get all categories: %s", err.Error())
		return nil, err
	}

	return categories, nil
}
