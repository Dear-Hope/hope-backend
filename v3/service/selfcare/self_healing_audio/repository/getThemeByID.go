package repository

import (
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"log"
)

func (ths *repository) GetThemeByID(id, userID uint) (*model.SelfHealingAudioTheme, error) {
	var selfHealingAudioTheme model.SelfHealingAudioTheme

	err := ths.db.Get(
		&selfHealingAudioTheme,
		`SELECT id, title FROM `+selfHealingAudioTheme.TableWithSchemaName()+` WHERE id = $1`,
		id,
	)
	if err != nil {
		log.Printf("get detail self healing audio theme: %s", err.Error())
		return nil, err
	}

	err = ths.db.Get(
		&selfHealingAudioTheme.LastPlayedOrder,
		`SELECT audio_order FROM "selfcare".self_healing_audio_histories WHERE theme_id = $1 AND user_id = $2 ORDER BY audio_order DESC LIMIT 1`,
		id,
		userID,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("get detail self healing audio theme order: %s", err.Error())
		return nil, err
	}

	var playlist model.SelfHealingAudios
	err = ths.db.Select(
		&playlist,
		`SELECT id, title, script_writer, image_url, duration, sha.order FROM `+model.SelfHealingAudio{}.TableWithSchemaName()+` sha WHERE theme_id = $1 ORDER BY sha.order`,
		id,
	)
	if err != nil {
		log.Printf("get detail self healing audio theme playlist: %s", err.Error())
		return nil, err
	}
	selfHealingAudioTheme.Playlist = playlist

	return &selfHealingAudioTheme, nil
}
