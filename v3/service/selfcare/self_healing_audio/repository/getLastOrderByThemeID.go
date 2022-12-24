package repository

import (
	"database/sql"
	"errors"
	"log"
)

func (ths *repository) GetLastOrderByThemeID(themeID, userID uint) (int, error) {
	var lastOrder int
	err := ths.db.Get(
		&lastOrder,
		`SELECT audio_order FROM "selfcare".self_healing_audio_histories WHERE theme_id = $1 AND user_id = $2 ORDER BY audio_order DESC LIMIT 1`,
		themeID, userID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		log.Printf("get last played self healing audio by theme id: %s", err.Error())
		return 0, err
	}

	return lastOrder, nil
}
