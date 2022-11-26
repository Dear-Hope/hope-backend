package repository

import (
	"log"
)

func (ths *repository) Delete(email string) error {
	_, err := ths.db.Queryx(
		`DELETE FROM "newsletter".subscriptions WHERE email = $1`,
		email,
	)
	if err != nil {
		log.Printf("delete subscription: %s", err.Error())
		return err
	}

	return nil
}
