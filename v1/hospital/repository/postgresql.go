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

func NewPostgreSQLRepository(db *gorm.DB) models.HospitalRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) GetAll(whereClause models.Hospital) ([]*models.Hospital, error) {
	var hospitals []*models.Hospital

	err := ths.db.Where(whereClause).Find(&hospitals).Error
	if err != nil {
		log.Printf("Hospital get all: %s", err.Error())

		err = errors.New("failed to get all Hospital")
		return hospitals, err
	}

	return hospitals, nil
}
