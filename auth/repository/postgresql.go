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

func (ths *postgreSQLRepository) CreateUser(user *models.User) error {
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

func (ths *postgreSQLRepository) CreateProfile(role string, profile models.Profile) error {
	var err error
	if role == "patient" {
		patient, _ := profile.(*models.Patient)
		err = ths.db.Create(patient).Error
	} else {
		psychologist, _ := profile.(*models.Psychologist)
		err = ths.db.Create(psychologist).Error
	}

	if err != nil {
		log.Printf("user profile create failed: %s", err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("this user already had a profile")
		}

		return errors.New("failed to create user profile")
	}

	return nil
}

func (ths *postgreSQLRepository) GetUserByEmail(whereClause *models.User) (*models.User, error) {
	var user models.User

	err := ths.db.Where(&whereClause).First(&user).Error
	if err != nil {
		log.Printf("user get by email: %s", err.Error())

		err = errors.New("user not found with given email")
		return nil, err
	}

	return &user, nil
}

func (ths *postgreSQLRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User

	err := ths.db.First(&user, id).Error
	if err != nil {
		log.Printf("user get by id: %s", err.Error())

		err = errors.New("user not found")
		return nil, err
	}

	return &user, nil
}

func (ths *postgreSQLRepository) GetProfileByUserID(role string, userID uint) (models.Profile, error) {
	var profile models.Profile
	if role == "patient" {
		var patient models.Patient
		err := ths.db.Where(&models.Patient{UserID: userID}).First(&patient).Error
		if err != nil {
			log.Printf("profile get by user id: %s", err.Error())

			err = errors.New("user profile not found")
			return nil, err
		}
		profile = &patient
	} else {
		var psychologist models.Psychologist
		err := ths.db.Where(&models.Psychologist{UserID: userID}).First(&psychologist).Error
		if err != nil {
			log.Printf("profile get by user id: %s", err.Error())

			err = errors.New("user profile not found")
			return nil, err
		}
		profile = &psychologist
	}

	return profile, nil
}

func (ths *postgreSQLRepository) UpdateUser(user *models.User) (*models.User, error) {
	var updatedUser models.User
	err := ths.db.Updates(user).Error
	if err != nil {
		log.Printf("user update: %s", err.Error())

		err = errors.New("failed to update user")
		return nil, err
	}

	ths.db.Where(user).First(&updatedUser)
	return &updatedUser, nil
}

func (ths *postgreSQLRepository) UpdateProfile(role string, profile models.Profile) (models.Profile, error) {
	var err error
	var updatedProfile models.Profile
	if role == "patient" {
		patient, _ := profile.(*models.Patient)
		err = ths.db.Where(models.Patient{UserID: patient.UserID}).Updates(patient).Error
		ths.db.Where(patient).First(patient)
		updatedProfile = patient
	} else {
		psychologist, _ := profile.(*models.Psychologist)
		err = ths.db.Where(models.Psychologist{UserID: psychologist.UserID}).Updates(psychologist).Error
		ths.db.Where(psychologist).First(psychologist)
		updatedProfile = psychologist
	}

	if err != nil {
		log.Printf("user profile update: %s", err.Error())

		err = errors.New("failed to update user profile")
		return nil, err
	}

	return updatedProfile, nil
}

func (ths *postgreSQLRepository) UpdateOnlineStatus(userID uint) error {
	var user models.User
	ths.db.Select("online_status").First(&user, userID)

	user.ID = userID
	err := ths.db.Model(&user).Update("online_status", !user.OnlineStatus).Error
	if err != nil {
		log.Printf("user online status update: %s", err.Error())

		err = errors.New("failed to update user online status")
		return err
	}

	return nil
}
