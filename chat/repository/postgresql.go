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

func chatOrderByDateCreated(db *gorm.DB) *gorm.DB {
	return db.Order("chat.created_at ASC")
}

func (ths *postgreSQLRepository) GetConversationByID(id uint) (*models.Conversation, error) {
	var conversation models.Conversation

	err := ths.db.Preload("Chats", chatOrderByDateCreated).Preload("FirstUser.Profile").Preload("SecondUser.Profile").First(&conversation, id).Error
	if err != nil {
		log.Printf("conversation get by id: %s", err.Error())

		err = errors.New("conversation not found")
		return nil, err
	}

	return &conversation, nil
}

func (ths *postgreSQLRepository) GetAllConversationByUserID(userID uint) ([]*models.Conversation, error) {
	var conversations []*models.Conversation

	err := ths.db.Preload("Chats", chatOrderByDateCreated).
		Preload("FirstUser.Profile").
		Preload("SecondUser.Profile").
		Where(models.Conversation{FirstUserID: userID}).
		Or(models.Conversation{SecondUserID: userID}).
		Find(&conversations).Error
	if err != nil {
		log.Printf("conversation get all by user id: %s", err.Error())

		err = errors.New("something wrong when get all conversation by user ID")
		return nil, err
	}

	return conversations, nil
}

func (ths *postgreSQLRepository) CreateChat(newChat models.Chat, conversation models.Conversation) (*models.Chat, error) {
	err := ths.db.Model(&conversation).Association("Chats").Append(&newChat)
	if err != nil {
		log.Printf("chat create failed: %s", err.Error())

		err = errors.New("failed to save chat")
		return nil, err
	}

	return &newChat, nil
}
