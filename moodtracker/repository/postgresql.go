package repository

import (
	"HOPE-backend/models"
	"errors"
	"log"
	"strings"

	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) models.MoodTrackerRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) Create(newEmotion models.Emotion) (*models.Emotion, error) {
	err := ths.db.Create(&newEmotion).Error
	if err != nil {
		log.Printf("new emotion create failed: %s", err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("this patient emotion for this time frame today is already exists")
		}

		return nil, errors.New("failed to create new emotion")
	}

	return &newEmotion, nil
}

func (ths *postgreSQLRepository) GetAllEmotionByPatientID(id uint) ([]*models.Emotion, error) {
	var emotions []*models.Emotion
	err := ths.db.Where(models.Emotion{PatientID: id}).Find(&emotions).Error
	if err != nil {
		log.Printf("emotions get all by patient id: %s", err.Error())

		err = errors.New("something wrong when get all emotions by patient ID")
		return nil, err
	}

	return emotions, nil
}
