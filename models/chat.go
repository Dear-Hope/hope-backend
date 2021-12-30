package models

import (
	"gorm.io/gorm"
)

type ChatType int

const (
	INVALID ChatType = iota
	TEXT
	IMAGE
)

var StringToChatType map[string]ChatType = map[string]ChatType{
	"INVALID": INVALID,
	"TEXT":    TEXT,
	"IMAGE":   IMAGE,
}

var ChatTypeToString map[ChatType]string = map[ChatType]string{
	INVALID: "INVALID",
	TEXT:    "TEXT",
	IMAGE:   "IMAGE",
}

type Chat struct {
	gorm.Model
	ConversationID uint     `json:"conversation_id"`
	User           User     `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID         uint     `json:"user_id"`
	Type           ChatType `json:"type"`
	Messages       string   `json:"message"`
}

type Conversation struct {
	gorm.Model
	FirstUser    User   `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
	SecondUser   User   `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
	FirstUserID  uint   `json:"first_user_id" gorm:"uniqueIndex:idx_first_second_user_id"`
	SecondUserID uint   `json:"second_user_id" gorm:"uniqueIndex:idx_first_second_user_id"`
	Chats        []Chat `json:"chats" `
}

func (Chat) TableName() string {
	return "chat"
}

func (Conversation) TableName() string {
	return "conversation"
}

type ChatService interface {
	NewConversation(NewConversationRequest) (*Conversation, error)
	GetConversation(uint) (*Conversation, error)
	ListConversation(uint) ([]*Conversation, error)
	NewChat(NewChatRequest) (*Chat, error)
}

type ChatRepository interface {
	CreateConversation(Conversation) (*Conversation, error)
	GetConversationByID(uint) (*Conversation, error)
	GetAllConversationByUserID(uint) ([]*Conversation, error)
	CreateChat(Chat, Conversation) (*Chat, error)
}

type NewConversationRequest struct {
	FirstUserID  uint `json:"first_user_id"`
	SecondUserID uint `json:"second_user_id"`
}

type NewChatRequest struct {
	ConversationID uint   `json:"conversation_id"`
	UserID         uint   `json:"user_id"`
	Type           string `json:"type"`
	Message        string `json:"message"`
}

type GetConversationResponse struct {
	ConversationID    uint        `json:"conversation_id"`
	FirstUserProfile  UserProfile `json:"first_user"`
	SecondUserProfile UserProfile `json:"second_user"`
	FirstUserID       uint        `json:"first_user_id"`
	SecondUserID      uint        `json:"second_user_id"`
	Chats             []Chat      `json:"chats"`
}
