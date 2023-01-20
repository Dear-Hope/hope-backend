package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) StoreHistory(newHistory model.BreathingExerciseHistory) (*model.BreathingExerciseHistory, error) {
	rows, err := ths.db.NamedQuery(
		`INSERT INTO `+newHistory.TableWithSchemaName()+` (exercise_id, user_id)
		VALUES (:exercise_id, :user_id)
		ON CONFLICT (exercise_id, user_id)
		DO UPDATE SET updated_at = now()
		RETURNING id`,
		newHistory,
	)
	if err != nil {
		log.Printf("new breathing exercise history create failed: %s", err.Error())
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		if err = rows.Scan(&newHistory.ID); err != nil {
			log.Printf("new breathing exercise history failed: %s", err.Error())
			return nil, err
		}
	}

	return &newHistory, nil
}
