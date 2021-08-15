package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Laboratory struct {
	gorm.Model
	Name        string         `json:"name" gorm:"unique;not null"`
	Location    string         `json:"location" gorm:"not null"`
	PhoneNumber string         `json:"phone_number" gorm:"not null"`
	About       string         `json:"about" gorm:"not null"`
	Address     string         `json:"address" gorm:"not null"`
	Services    pq.StringArray `json:"services" gorm:"not null;type:text[]"`
	Image       string         `json:"image"`
}

func (Laboratory) TableName() string {
	return "laboratory"
}

type LaboratoryService interface {
	List(string, string, string) ([]*Laboratory, error)
}

type LaboratoryRepository interface {
	GetAll(Laboratory) ([]*Laboratory, error)
}
