package models

import (
	"gorm.io/gorm"
)

type ChatType int

const (
	Text ChatType = iota + 1
	Image
)

type Chat struct {
	gorm.Model
	ConversationID int      `json:"conversation_id"`
	User           User     `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID         int      `json:"user_id"`
	Type           ChatType `json:"type"`
	Messages       string   `json:"message"`
}

type Conversation struct {
	gorm.Model
	FirstUser    User   `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
	SecondUser   User   `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
	FirstUserID  int    `json:"first_user_id" gorm:"uniqueIndex:idx_first_second_user_id"`
	SecondUserID int    `json:"second_user_id" gorm:"uniqueIndex:idx_first_second_user_id"`
	Chats        []Chat `json:"chats"`
}

func (Chat) TableName() string {
	return "chat"
}

func (Conversation) TableName() string {
	return "conversation"
}

type ChatService interface {
	NewConversation(NewConversationRequest) (*Conversation, error)
}

type ChatRepository interface {
	CreateConversation(Conversation) (*Conversation, error)
}

type NewConversationRequest struct {
	FirstUserID  uint `json:"first_user_id"`
	SecondUserID uint `json:"second_user_id"`
}
