package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) Store(newEmotion model.Emotion) (*model.Emotion, error) {
	rows, err := ths.db.NamedQuery(
		`WITH rows AS (
			INSERT INTO "moodtracker".emotions (mood_id, description, time_frame, date, user_id, triggers, feelings, scale)
			VALUES (:mood_id, :description, :time_frame, :date, :user_id, :triggers, :feelings, :scale)
			RETURNING id,mood_id
		)
		SELECT r.id as id, m.name as mood
		FROM rows r, "moodtracker".moods m
		WHERE r.mood_id = m.id`,
		newEmotion,
	)
	if err != nil {
		log.Printf("new emotion create failed: %s", err.Error())
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		if err = rows.Scan(&newEmotion.ID, &newEmotion.Mood); err != nil {
			log.Printf("new emotion create failed: %s", err.Error())
			return nil, err
		}
	}

	return &newEmotion, nil
}
