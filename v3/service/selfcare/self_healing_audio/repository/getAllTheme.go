package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAllTheme() (model.SelfHealingAudioThemes, error) {
	var selfHealingAudioThemes model.SelfHealingAudioThemes

	query := `SELECT id, title, description, image_url FROM ` + model.SelfHealingAudioTheme{}.TableWithSchemaName() + ` ORDER BY id`

	err := ths.db.Select(&selfHealingAudioThemes, query)
	if err != nil {
		log.Printf("get all self healing audio themes: %s", err.Error())
		return nil, err
	}

	for i := range selfHealingAudioThemes {
		err := ths.db.Get(&selfHealingAudioThemes[i].TotalAudio, `SELECT COUNT(*) AS total_audio FROM `+model.SelfHealingAudio{}.TableWithSchemaName()+` WHERE theme_id = $1`, selfHealingAudioThemes[i].ID)
		if err != nil {
			log.Printf("get all self healing audio theme playlist count: %s", err.Error())
			return nil, err
		}
	}

	return selfHealingAudioThemes, nil
}
