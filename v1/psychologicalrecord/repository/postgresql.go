package repository

import (
	"HOPE-backend/v1/models"
	"errors"
	"log"
	"strings"

	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) models.PsychologicalRecordRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) Create(newRecord models.PsychologicalRecord) (*models.PsychologicalRecord, error) {
	err := ths.db.Create(&newRecord).Error
	if err != nil {
		log.Printf("psychological record create failed: %s", err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("record for this patient is already exists")
		}

		return nil, errors.New("failed to create psychological record")
	}

	return &newRecord, nil
}

func (ths *postgreSQLRepository) GetRecordByID(id uint) (*models.PsychologicalRecord, error) {
	var record models.PsychologicalRecord
	err := ths.db.First(&record, id).Error
	if err != nil {
		log.Printf("psychological record get by id: %s", err.Error())

		err = errors.New("psychological record not found")
		return nil, err
	}

	return &record, nil
}

func (ths *postgreSQLRepository) GetAllRecordByPyschologistID(pyschologistID uint) (
	[]*models.PsychologicalRecord,
	error,
) {
	var records []*models.PsychologicalRecord

	err := ths.db.Where(models.PsychologicalRecord{PsychologistID: pyschologistID}).
		Find(&records).
		Error
	if err != nil {
		log.Printf("psychological record get all by psychologist id: %s", err.Error())

		err = errors.New("something wrong when get all record by psychologist ID")
		return nil, err
	}

	return records, nil
}
