package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAllAudioByMoodID(moodID uint) (model.SelfHealingAudios, error) {
	var (
		selfHealingAudios model.SelfHealingAudios
		where             string
		args              []interface{}
	)

	if moodID > 0 {
		where = `WHERE mood_id = $1`
		args = append(args, moodID)
	}

	err := ths.db.Select(
		&selfHealingAudios,
		`SELECT id, theme_id, title, image_url, audio_url, description, benefit, voice_over, script_writer, sha.order FROM `+model.SelfHealingAudio{}.TableWithSchemaName()+` sha `+where,
		args...,
	)
	if err != nil {
		log.Printf("get all self healing audio by mood id: %s", err.Error())
		return nil, err
	}

	return selfHealingAudios, nil
}
