package chat

import (
	"HOPE-backend/models"
	"errors"
	"fmt"
)

type service struct {
	chatRepo models.ChatRepository
	userRepo models.AuthRepository
}

func NewChatService(chatRepo models.ChatRepository, userRepo models.AuthRepository) models.ChatService {
	return &service{
		chatRepo: chatRepo,
		userRepo: userRepo,
	}
}

func (ths *service) NewConversation(req models.NewConversationRequest) (*models.Conversation, error) {
	if req.FirstUserID > req.SecondUserID {
		req.FirstUserID, req.SecondUserID = req.SecondUserID, req.FirstUserID
	}

	firstUser, err := ths.userRepo.GetByID(req.FirstUserID)
	if err != nil {
		return nil, err
	}

	secondUser, err := ths.userRepo.GetByID(req.SecondUserID)
	if err != nil {
		return nil, err
	}

	newConversation := models.Conversation{
		FirstUser:  *firstUser,
		SecondUser: *secondUser,
	}

	conversation, err := ths.chatRepo.CreateConversation(newConversation)

	return conversation, err
}

func (ths *service) GetConversation(id uint) (*models.Conversation, error) {
	conversation, err := ths.chatRepo.GetConversationByID(id)
	if err != nil {
		return nil, err
	}

	return conversation, nil
}

func (ths *service) ListConversation(userID uint) ([]*models.Conversation, error) {
	conversations, err := ths.chatRepo.GetAllConversationByUserID(userID)
	if err != nil {
		return nil, err
	}

	return conversations, nil
}

func (ths *service) NewChat(req models.NewChatRequest) (*models.Chat, error) {
	conversation, err := ths.chatRepo.GetConversationByID(req.ConversationID)
	if err != nil {
		return nil, err
	}

	if req.UserID != conversation.FirstUserID && req.UserID != conversation.SecondUserID {
		return nil, errors.New("user is not one of the conversation owners")
	}

	newChat := models.Chat{
		UserID:   req.UserID,
		Type:     models.StringToChatType[req.Type],
		Messages: req.Message,
	}

	fmt.Println(newChat)

	chat, err := ths.chatRepo.CreateChat(newChat, *conversation)
	if err != nil {
		return nil, err
	}

	return chat, nil
}
