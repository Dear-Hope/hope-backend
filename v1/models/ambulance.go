package models

import "gorm.io/gorm"

type Ambulance struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"`
	Location    string `json:"location" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
	Address     string `json:"address" gorm:"not null"`
	Price       int    `json:"price" gorm:"not null"`
}

func (Ambulance) TableName() string {
	return "ambulance"
}

type AmbulanceService interface {
	List(string, string) ([]*Ambulance, error)
}

type AmbulanceRepository interface {
	GetAll(Ambulance) ([]*Ambulance, error)
}
