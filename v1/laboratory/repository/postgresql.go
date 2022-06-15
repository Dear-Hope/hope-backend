package repository

import (
	"HOPE-backend/v1/models"
	"errors"
	"log"

	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) models.LaboratoryRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) GetAll(whereClause models.Laboratory) ([]*models.Laboratory, error) {
	var laboratories []*models.Laboratory

	err := ths.db.Where(whereClause).Find(&laboratories).Error
	if err != nil {
		log.Printf("Laboratory get all: %s", err.Error())

		err = errors.New("failed to get all Laboratory")
		return laboratories, err
	}

	return laboratories, nil
}
