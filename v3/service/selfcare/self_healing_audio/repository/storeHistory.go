package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) StoreHistory(newAudio model.SelfHealingAudioHistory) (*model.SelfHealingAudioHistory, error) {
	rows, err := ths.db.NamedQuery(
		`INSERT INTO `+newAudio.TableWithSchemaName()+` (theme_id, audio_id, user_id, audio_order)
		VALUES (:theme_id, :audio_id, :user_id, :audio_order)
		ON CONFLICT (theme_id, audio_id, user_id)
		DO UPDATE SET updated_at = now()
		RETURNING id`,
		newAudio,
	)
	if err != nil {
		log.Printf("new self healing audio history create failed: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&newAudio.ID); err != nil {
			log.Printf("new self healing audio history failed: %s", err.Error())
			return nil, err
		}
	}

	return &newAudio, nil
}
