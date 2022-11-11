package auth

import "HOPE-backend/v3/model"

type Repository interface {
	Create(model.User) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	GetByID(uint) (*model.User, error)
	Update(model.User) (*model.User, error)
	SetToActive(uint) error
	SetProfilePhoto(uint, string) error
	UpdatePassword(uint, string) error
	DeleteByEmail(string) error
}
