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

	err := ths.db.Joins("Profile").Where(&whereClause).First(&user).Error
	if err != nil {
		log.Printf("user get by email: %s", err.Error())

		err = errors.New("user not found with given email")
		return nil, err
	}

	return &user, nil
}

func (ths *postgreSQLRepository) GetByID(id uint) (*models.User, error) {
	var user models.User

	err := ths.db.Joins("Profile").First(&user, id).Error
	if err != nil {
		log.Printf("user get by id: %s", err.Error())

		err = errors.New("user not found")
		return nil, err
	}

	return &user, nil
}

func (ths *postgreSQLRepository) Update(user *models.User) (*models.User, error) {
	err := ths.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(user).Error
	if err != nil {
		log.Printf("user update: %s", err.Error())

		err = errors.New("failed to update user")
		return nil, err
	}

	var updatedUser models.User
	ths.db.Joins("Profile").First(&updatedUser, user.ID)

	return &updatedUser, nil
}
