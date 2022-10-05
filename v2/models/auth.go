package models

import "mime/multipart"

type DBUserWithProfile struct {
	UserID       uint   `db:"user_id"`
	ProfileID    uint   `db:"profile_id"`
	Email        string `db:"email"`
	Password     string `db:"password"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	ProfilePhoto string `db:"photo"`
	Job          string `db:"job"`
	Activities   string `db:"activities"`
	IsActive     bool   `db:"is_active"`
	SecretKey    string `db:"secret_key"`
}

type User struct {
	ID        uint   `json:"id" `
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsActive  bool   `json:"is_active"`
	SecretKey string `json:"secret_key"`
}

type Profile struct {
	ID         uint   `json:"id"`
	Job        string `json:"job"`
	Activities string `json:"activities"`
	Photo      string `json:"photo"`
}

type AuthService interface {
	Login(LoginRequest) (*TokenPair, error)
	Register(RegisterRequest) (*TokenPair, error)
	GetLoggedInUser(uint) (*UserResponse, error)
	UpdateLoggedInUser(UpdateRequest) (*UserResponse, error)
	Activate(ActivateRequest) (*TokenPair, error)
	ResetPassword(ResetPasswordRequest) error
	ChangePassword(ChangePasswordRequest) (*TokenPair, error)
	SaveProfilePhoto(SaveProfilePhotoRequest) (string, error)
	ResendActivationCode(ResetPasswordRequest) error
	DeleteUser(ResetPasswordRequest) error
}

type AuthRepository interface {
	CreateUserWithProfile(DBUserWithProfile) (uint, uint, error)
	GetUserWithProfileByEmail(string) (*DBUserWithProfile, error)
	GetUserWithProfileByID(uint) (*DBUserWithProfile, error)
	UpdateUserWithProfile(DBUserWithProfile) (*DBUserWithProfile, error)
	SetUserToActive(uint) error
	SetUserProfilePhoto(uint, string) error
	UpdatePassword(uint, string) error
	DeleteUserByEmail(string) error
}

type UserResponse struct {
	User
	Profile Profile `json:"profile,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email        string               `json:"email"`
	Password     string               `json:"password"`
	FirstName    string               `json:"first_name,omitempty"`
	LastName     string               `json:"last_name,omitempty"`
	ProfilePhoto string               `json:"profile_photo,omitempty"`
	Profile      CreateProfileRequest `json:"profile,omitempty"`
}

type CreateProfileRequest struct {
	Weight     float32 `json:"weight"`
	Height     float32 `json:"height"`
	Job        string  `json:"job,omitempty"`
	Activities string  `json:"activities,omitempty"`
}

type UpdateRequest struct {
	RegisterRequest
	UserID    uint `json:"user_id,omitempty"`
	ProfileID uint `json:"profile_id,omitempty"`
}

type TokenPair struct {
	Access  string `json:"access,omitempty"`
	Refresh string `json:"refresh,omitempty"`
}

type ActivateRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type ChangePasswordRequest struct {
	Key         string `json:"key"`
	NewPassword string `json:"new_password"`
}

type EmailTemplate struct {
	Content string
	Email   string
	Subject string
}

type SaveProfilePhotoRequest struct {
	File      *multipart.File
	Extension string
	UserID    uint
}
