package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Hospital struct {
	gorm.Model
	Name        string         `json:"name" gorm:"unique;not null"`
	Location    string         `json:"location" gorm:"not null"`
	PhoneNumber string         `json:"phone_number" gorm:"not null"`
	About       string         `json:"about" gorm:"not null"`
	Address     string         `json:"address" gorm:"not null"`
	Polyclinics pq.StringArray `json:"polyclinics" gorm:"not null;type:text[]"`
	Image       string         `json:"image"`
}

func (Hospital) TableName() string {
	return "hospital"
}

type HospitalService interface {
	List(string, string, string) ([]*Hospital, error)
}

type HospitalRepository interface {
	GetAll(Hospital) ([]*Hospital, error)
}
