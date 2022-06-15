package repository

import (
	"HOPE-backend/v1/models"
	"errors"
	"log"
	"strings"

	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) models.SelfCareRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) Create(newItem models.SelfCareItem) (*models.SelfCareItem, error) {
	err := ths.db.Create(&newItem).Error
	if err != nil {
		log.Printf("new self care item create failed: %s", err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("this self care item already exist")
		}

		return nil, errors.New("failed to create new self care item")
	}

	return &newItem, nil
}

func (ths *postgreSQLRepository) GetItemsByMood(mood string) ([]*models.SelfCareItem, error) {
	var items []*models.SelfCareItem
	err := ths.db.Where("mood = ?", mood).Find(&items).Error
	if err != nil {
		log.Printf("self care items get by mood: %s", err.Error())

		err = errors.New("something wrong when get all self care items by mood")
		return nil, err
	}

	return items, nil
}

func (ths *postgreSQLRepository) GetAllItems() ([]*models.SelfCareItem, error) {
	var items []*models.SelfCareItem
	err := ths.db.Find(&items).Error
	if err != nil {
		log.Printf("self care items get all: %s", err.Error())

		err = errors.New("something wrong when get all self care items")
		return nil, err
	}

	return items, nil
}
