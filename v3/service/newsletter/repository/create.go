package repository

import (
	"HOPE-backend/v3/model"
	"log"
	"strings"
)

func (ths *repository) Create(newSubscription model.Subscription) error {
	_, err := ths.db.NamedQuery(
		`INSERT INTO "newsletter".subscriptions (id, email, subscribed_at) VALUES (:id, :email, :subscribed_at)`,
		&newSubscription,
	)
	if err != nil {
		log.Printf("new subscription create failed: %s", err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return nil
		}

		return err
	}

	return nil
}
