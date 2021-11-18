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

func NewPostgreSQLRepository(db *gorm.DB) models.ChatRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (ths *postgreSQLRepository) CreateConversation(newConversation models.Conversation) (*models.Conversation, error) {
	err := ths.db.Create(&newConversation).Error
	if err != nil {
		log.Printf("conversation create failed: %s", err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("conversation betweetn this two user already exists")
		}

		return nil, errors.New("failed to create conversation")
	}

	return &newConversation, nil
}
