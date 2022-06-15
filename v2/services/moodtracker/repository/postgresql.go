package repository

import (
	"HOPE-backend/v2/models"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type postgreSQLRepository struct {
	db *sqlx.DB
}

func NewPostgreSQLRepository(db *sqlx.DB) models.MoodTrackerRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) Create(newEmotion models.Emotion) (*models.Emotion, error) {
	rows, err := ths.db.NamedQuery(
		`INSERT INTO "moodtracker".emotions (mood, description, time_frame, date, user_id, triggers) 
		VALUES (:mood, :description, :time_frame, :date, :user_id, :triggers) 
		RETURNING id`,
		newEmotion,
	)
	if err != nil {
		log.Printf("new emotion create failed: %s", err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("this patient emotion for this time frame today is already exists")
		}

		return nil, errors.New("failed to create new emotion")
	}

	for rows.Next() {
		if err = rows.Scan(&newEmotion.ID); err != nil {
			log.Printf("new emotion create failed: %s", err.Error())
			return nil, errors.New("failed to create new emotion")
		}
	}

	return &newEmotion, nil
}

func (ths *postgreSQLRepository) GetAllEmotionByUserID(id uint) ([]*models.Emotion, error) {
	emotions := []*models.Emotion{}
	err := ths.db.Select(
		&emotions,
		`SELECT id, mood, description, time_frame, date, triggers, user_id 
		FROM "moodtracker".emotions WHERE user_id = $1`,
		id,
	)
	if err != nil {
		log.Printf("emotions get all by patient id: %s", err.Error())

		err = errors.New("something wrong when get all emotions by patient ID")
		return nil, err
	}

	return emotions, nil
}

func (ths *postgreSQLRepository) GetAllEmotionByUserIDPerWeek(id uint) ([]*models.Emotion, error) {
	emotions := []*models.Emotion{}
	err := ths.db.Select(
		&emotions,
		`SELECT id, mood, description, time_frame, date, triggers, user_id 
		FROM "moodtracker".emotions 
		WHERE user_id = $1 AND date >= $2 AND date <= $3`,
		id, getStartDayOfWeek(), getLastDayOfWeek(),
	)
	if err != nil {
		log.Printf("emotions get all by patient id per week: %s", err.Error())

		err = errors.New("something wrong when get all emotions by patient ID per week")
		return nil, err
	}

	return emotions, nil
}

func (ths *postgreSQLRepository) GetAllEmotionByUserIDPerMonth(id uint) ([]*models.Emotion, error) {
	emotions := []*models.Emotion{}
	firstDay, lastDay := getFirstAndLastDayOfMonth()
	err := ths.db.Select(
		&emotions,
		`SELECT id, mood, description, time_frame, date, triggers, user_id 
		FROM "moodtracker".emotions 
		WHERE user_id = $1 AND date >= $2 AND date <= $3`,
		id, firstDay, lastDay,
	)
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
