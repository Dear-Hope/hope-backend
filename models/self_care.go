package models

import (
	"gorm.io/gorm"
)

type SelfCareItem struct {
	gorm.Model
	Link        string `json:"link" gorm:"not null"`
	Mood        string `json:"mood" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Type        string `json:"type" gorm:"not null"`
	Description string `json:"description"`
}

func (SelfCareItem) TableName() string {
	return "self_care"
}

type SelfCareService interface {
	NewItem(NewSelfCareItemRequest) (*SelfCareItem, error)
	GetItemsByMood(string) ([]*SelfCareItem, error)
}

type SelfCareRepository interface {
	Create(SelfCareItem) (*SelfCareItem, error)
	GetItemsByMood(string) ([]*SelfCareItem, error)
}

type NewSelfCareItemRequest struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Mood        string `json:"mood"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}
