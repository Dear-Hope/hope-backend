package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) Store(newEmotion model.Emotion) (*model.Emotion, error) {
	rows, err := ths.db.NamedQuery(
		`INSERT INTO "moodtracker".emotions (mood, description, time_frame, date, user_id, triggers, feelings, scale)
		VALUES (:mood, :description, :time_frame, :date, :user_id, :triggers, :feelings, :scale)
		RETURNING id`,
		newEmotion,
	)
	if err != nil {
		log.Printf("new emotion create failed: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&newEmotion.ID); err != nil {
			log.Printf("new emotion create failed: %s", err.Error())
			return nil, err
		}
	}

	return &newEmotion, nil
}
