package repository

import (
	"HOPE-backend/models"
	"errors"
	"log"

	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) models.AmbulanceRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) GetAll(whereClause models.Ambulance) ([]*models.Ambulance, error) {
	var ambulances []*models.Ambulance

	err := ths.db.Where(whereClause).Find(&ambulances).Error
	if err != nil {
		log.Printf("Ambulance get all: %s", err.Error())

		err = errors.New("failed to get all Ambulance")
		return ambulances, err
	}

	return ambulances, nil
}
