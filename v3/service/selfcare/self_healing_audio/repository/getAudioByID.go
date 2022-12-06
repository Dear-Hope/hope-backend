package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAudioByID(id uint) (*model.SelfHealingAudio, error) {
	var selfHealingAudio model.SelfHealingAudio

	err := ths.db.Get(
		&selfHealingAudio,
		`SELECT id, theme_id, title, image_url, audio_url, description, benefit, voice_over, script_writer, sha.order FROM `+selfHealingAudio.TableWithSchemaName()+` sha WHERE id = $1`,
		id,
	)
	if err != nil {
		log.Printf("get detail self healing audio: %s", err.Error())
		return nil, err
	}

	var subtitles model.SelfHealingAudioSubtitles
	err = ths.db.Select(
		&subtitles,
		`SELECT id, text, start FROM `+model.SelfHealingAudioSubtitle{}.TableWithSchemaName()+` shas WHERE audio_id = $1 ORDER BY shas.order`,
		id,
	)
	if err != nil {
		log.Printf("get detail self healing audio subtitles: %s", err.Error())
		return nil, err
	}
	selfHealingAudio.Subtitles = subtitles

	return &selfHealingAudio, nil
}
