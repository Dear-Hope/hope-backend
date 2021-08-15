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

func NewPostgreSQLRepository(db *gorm.DB) models.AuthRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) Create(user *models.User) error {
	err := ths.db.Create(user).Error
	if err != nil {
		log.Printf("user create failed: %s", err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("user with this email address already exists")
		}

		return errors.New("failed to create user")
	}

	return nil
}

func (ths *postgreSQLRepository) GetByEmail(whereClause *models.User) (*models.User, error) {
	var user models.User

	err := ths.db.Where(&whereClause).First(&user).Error
	if err != nil {
		log.Printf("user get by email: %s", err.Error())

		err = errors.New("user not found with given email")
	}

	return &user, err
}
