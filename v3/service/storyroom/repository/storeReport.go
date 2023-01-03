package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) StoreReport(newReport model.Report) (*model.Report, error) {
	rows, err := ths.db.NamedQuery(
		`WITH rows AS (
			INSERT INTO "storyroom".reports (user_id, post_id, reason_id) 
			VALUES (:user_id, :post_id, :reason_id) 
			RETURNING id, reason_id
		)
		SELECT r.id, rr.reason FROM rows r, "storyroom".report_reasons rr
		WHERE r.reason_id = rr.id
		`,
		newReport,
	)
	if err != nil {
		log.Printf("new post report failed: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&newReport.ID, &newReport.Reason); err != nil {
			log.Printf("new post report failed: %s", err.Error())
			return nil, err
		}
	}

	return &newReport, nil
}
