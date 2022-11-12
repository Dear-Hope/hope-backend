package repository

import (
	"errors"
	"log"
)

func (ths *repository) DeleteByEmail(email string) error {
	_, err := ths.db.Queryx(
		`DELETE FROM "auth".users WHERE email = $1`,
		email,
	)
	if err != nil {
		log.Printf("delete user: %s", err.Error())
		return errors.New("failed to delete user")
	}

	return nil
}
