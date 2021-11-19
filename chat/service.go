package chat

import (
	"HOPE-backend/models"
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
