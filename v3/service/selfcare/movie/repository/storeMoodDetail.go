package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) StoreMoodDetail(movieID uint, moodIDs []uint) (model.Moods, error) {
	var (
		mood  model.Mood
		moods model.Moods
	)

	for _, moodID := range moodIDs {
		rows, err := ths.db.NamedQuery(
			`WITH rows AS (
				INSERT INTO "selfcare".movie_mood_details (movie_id, mood_id)
				VALUES (:movie_id, :mood_id)
				RETURNING mood_id
			)
			SELECT m.id as id, m.name as mood
			FROM rows r, "moodtracker".moods m
			WHERE r.mood_id = m.id`,
			map[string]interface{}{"movie_id": movieID, "mood_id": moodID},
		)
		if err != nil {
			log.Printf("new movie mood detail create failed: %s", err.Error())
			return nil, err
		}

		for rows.Next() {
			if err = rows.Scan(&mood.ID, &mood.Name); err != nil {
				log.Printf("new movie mood detail create failed: %s", err.Error())
				return nil, err
			}
		}

		moods = append(moods, mood)
		_ = rows.Close()
	}

	return moods, nil
}
