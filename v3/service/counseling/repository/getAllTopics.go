package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAllTopics() (model.Topics, error) {
	var topics model.Topics

	err := ths.db.Select(
		&topics,
		`SELECT id, name, image_url FROM "counseling".topics`,
	)
	if err != nil {
		log.Printf("get all topics: %s", err.Error())
		return nil, err
	}

	return topics, nil
}
