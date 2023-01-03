package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAllReason() (model.ReportReasons, error) {
	var reason model.ReportReasons

	err := ths.db.Select(
		&reason,
		`SELECT id, reason FROM "storyroom".report_reasons`,
	)
	if err != nil {
		log.Printf("get all reason: %s", err.Error())
		return nil, err
	}

	return reason, nil
}
