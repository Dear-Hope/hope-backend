package repository

import (
	"HOPE-backend/v1/models"
	"errors"
	"log"
	"strings"
	"time"

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

func (ths *postgreSQLRepository) GetAllEmotionByPatientIDPerWeek(id uint) ([]*models.Emotion, error) {
	var emotions []*models.Emotion
	err := ths.db.Where(
		"patient_id = ? AND date >= ? AND date <= ?",
		id,
		getStartDayOfWeek(),
		getLastDayOfWeek(),
	).Find(&emotions).Error
	if err != nil {
		log.Printf("emotions get all by patient id per week: %s", err.Error())

		err = errors.New("something wrong when get all emotions by patient ID per week")
		return nil, err
	}

	return emotions, nil
}

func (ths *postgreSQLRepository) GetAllEmotionByPatientIDPerMonth(id uint) ([]*models.Emotion, error) {
	var emotions []*models.Emotion
	firstDay, lastDay := getFirstAndLastDayOfMonth()
	err := ths.db.Where(
		"patient_id = ? AND date >= ? AND date <= ?",
		id,
		firstDay,
		lastDay,
	).Find(&emotions).Error
	if err != nil {
		log.Printf("emotions get all by patient id per week: %s", err.Error())

		err = errors.New("something wrong when get all emotions by patient ID per week")
		return nil, err
	}

	return emotions, nil
}

func getStartDayOfWeek() int64 { //get monday 00:00:00
	tm := time.Now()
	weekday := time.Duration(tm.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := tm.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return currentZeroDay.Add(-1 * (weekday - 1) * 24 * time.Hour).UnixMilli()
}

func getFirstAndLastDayOfMonth() (int64, int64) { //get first day of the month
	t := time.Now()
	firstday := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	lastday := firstday.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

	return firstday.UnixMilli(), lastday.UnixMilli()
}

func getLastDayOfWeek() int64 { //get sunday 00:00:00
	tm := time.Now()
	weekday := time.Duration(tm.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := tm.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return currentZeroDay.Add(1 * (7 - weekday) * 24 * time.Hour).UnixMilli()
}