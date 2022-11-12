package repository

import "log"

func (ths *repository) UpdatePassword(id uint, newPassword string) error {
	_, err := ths.db.Queryx(
		`UPDATE "auth".users SET password = $1 WHERE id = $2`,
		newPassword,
		id,
	)
	if err != nil {
		log.Printf("user update password: %s", err.Error())
		return err
	}

	return nil
}
