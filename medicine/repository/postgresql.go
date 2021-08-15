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

func NewPostgreSQLRepository(db *gorm.DB) models.MedicineRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) GetAll(whereClause models.Medicine) ([]*models.Medicine, error) {
	var medicines []*models.Medicine

	err := ths.db.Where(&whereClause).Find(&medicines).Error
	if err != nil {
		log.Printf("medicine get all: %s", err.Error())

		err = errors.New("failed to get all medicine")
		return medicines, err
	}

	return medicines, nil
}
