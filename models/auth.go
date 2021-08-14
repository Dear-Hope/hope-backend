package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string      `json:"email" gorm:"unique;not null"`
	Password string      `json:"password" gorm:"not null"`
	IsAdmin  bool        `json:"is_admin" gorm:"not null"`
	Profile  UserProfile `json:"profile"`
}

type UserProfile struct {
	gorm.Model
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Weight         float32 `json:"weight" gorm:"not null"`
	Height         float32 `json:"height" gorm:"not null"`
	Job            string  `json:"job"`
	Activities     string  `json:"activities"`
	DiseaseHistory string  `json:"disease_history"`
	UserID         uint    `json:"user_id"`
}

func (User) TableName() string {
	return "user"
}

func (UserProfile) TableName() string {
	return "user_profile"
}

type AuthService interface {
	Login(LoginRequest) (*TokenPair, error)
	Register(RegisterRequest) (*TokenPair, error)
}

type AuthRepository interface {
	Create(*User) error
	GetByEmailAndPassword(*User) (*User, error)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	FirstName      string  `json:"first_name,omitempty"`
	LastName       string  `json:"last_name,omitempty"`
	Weight         float32 `json:"weight"`
	Height         float32 `json:"height"`
	Job            string  `json:"job,omitempty"`
	Activities     string  `json:"activities,omitempty"`
	DiseaseHistory string  `json:"disease_history,omitempty"`
}

type TokenPair struct {
	Access  string `json:"access,omitempty"`
	Refresh string `json:"refresh,omitempty"`
}
