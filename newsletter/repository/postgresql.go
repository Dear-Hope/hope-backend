package repository

import (
	"HOPE-backend/models"
	"errors"
	"log"
	"strings"

	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) models.NewsletterRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) Create(newSubscription models.Subscription) error {
	err := ths.db.Create(&newSubscription).Error
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
	err := ths.db.Delete(models.Subscription{Email: email}).Error
	if err != nil {
		log.Printf("delete subscription: %s", err.Error())

		err = errors.New("something wrong when deleting subcription with email: " + email)
		return err
	}

	return nil
}
