package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetLastAudio(userID uint) (*model.SelfHealingAudio, error) {
	var selfHealingAudio model.SelfHealingAudio

	err := ths.db.Get(
		&selfHealingAudio,
		`WITH audio AS (
			SELECT audio_id AS id FROM `+model.SelfHealingAudioHistory{}.TableWithSchemaName()+`
			WHERE user_id = $1 ORDER BY updated_at DESC LIMIT 1
		)
		SELECT id, title, image_url, audio_url, description, benefit, voice_over, script_writer FROM `+selfHealingAudio.TableWithSchemaName()+` WHERE id = (SELECT id FROM audio)
		`,
		userID,
	)
	if err != nil {
		log.Printf("get last played self healing audio: %s", err.Error())
		return nil, err
	}

	var subtitles model.SelfHealingAudioSubtitles
	err = ths.db.Select(&subtitles, `SELECT id, text, start FROM `+model.SelfHealingAudioSubtitle{}.TableWithSchemaName()+` shas WHERE audio_id = $1 ORDER BY shas.order`, selfHealingAudio.ID)
	if err != nil {
		log.Printf("get last played self healing audio subtitles: %s", err.Error())
		return nil, err
	}

	selfHealingAudio.Subtitles = subtitles

	return &selfHealingAudio, nil
}
