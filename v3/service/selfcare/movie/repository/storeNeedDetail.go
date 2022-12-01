package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) StoreNeedDetail(movieID uint, needIDs []uint) (model.Needs, error) {
	var (
		need  model.Need
		needs model.Needs
	)

	for _, needID := range needIDs {
		rows, err := ths.db.NamedQuery(
			`WITH rows AS (
				INSERT INTO "selfcare".movie_need_details (movie_id, need_id)
				VALUES (:movie_id, :need_id)
				RETURNING need_id
			)
			SELECT n.id as id, n.name as need
			FROM rows r, "selfcare".needs n
			WHERE r.need_id = n.id`,
			map[string]interface{}{"movie_id": movieID, "need_id": needID},
		)
		if err != nil {
			log.Printf("new movie need detail create failed: %s", err.Error())
			return nil, err
		}

		for rows.Next() {
			if err = rows.Scan(&need.ID, &need.Name); err != nil {
				log.Printf("new movie need detail create failed: %s", err.Error())
				return nil, err
			}
		}

		needs = append(needs, need)
	}

	return needs, nil
}
