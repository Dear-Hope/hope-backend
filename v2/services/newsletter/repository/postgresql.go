package repository

import (
	"HOPE-backend/v2/models"
	"errors"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

type postgreSQLRepository struct {
	db *sqlx.DB
}

func NewPostgreSQLRepository(db *sqlx.DB) models.NewsletterRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) Create(newSubscription models.Subscription) error {
	_, err := ths.db.NamedQuery(
		`INSERT INTO "newsletter".subscriptions (id, email, subscribed_at) VALUES (:id, :email, :subscribed_at)`,
		&newSubscription,
	)
	if err != nil {
		log.Printf("new subscription create failed: %s", err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return nil
		}

		return errors.New("failed to create new subscription")
	}

	return nil
}

func (ths *postgreSQLRepository) Delete(email string) error {
	_, err := ths.db.Queryx(
		`DELETE FROM "newsletter".subscriptions WHERE email = $1`,
		email,
	)
	if err != nil {
		log.Printf("delete subscription: %s", err.Error())

		return errors.New("something wrong when deleting subcription with email: " + email)
	}

	return nil
}
