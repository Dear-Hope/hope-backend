package repository

import "log"

func (ths *repository) SetToActive(id uint) error {
	_, err := ths.db.Queryx(
		`UPDATE "auth".users SET is_active = true WHERE id = $1`,
		id,
	)
	if err != nil {
		log.Printf("user set to active: %s", err.Error())
		return err
	}

	return nil
}
