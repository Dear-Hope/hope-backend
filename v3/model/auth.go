package model

import "mime/multipart"

// import "mime/multipart"

type (
	User struct {
		UserID    uint   `db:"user_id"`
		Email     string `db:"email"`
		Password  string `db:"password"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		IsActive  bool   `db:"is_active"`
		SecretKey string `db:"secret_key"`

		Profile
	}

	Profile struct {
		ProfileID    uint   `db:"profile_id"`
		ProfilePhoto string `db:"photo"`
		Job          string `db:"job"`
		Activities   string `db:"activities"`
	}
)

func (ths User) TableWithSchemaName() string {
	return `"auth".users`
}

func (ths Profile) TableWithSchemaName() string {
	return `"auth".profiles`
}

type (
	UserResponse struct {
		ID        uint            `json:"id" `
		Email     string          `json:"email"`
		FirstName string          `json:"first_name"`
		LastName  string          `json:"last_name"`
		Profile   ProfileResponse `json:"profile,omitempty"`
	}

	ProfileResponse struct {
		ID         uint   `json:"id"`
		Job        string `json:"job"`
		Activities string `json:"activities"`
		Photo      string `json:"photo"`
	}
)

func (ths User) ToUserResponse() *UserResponse {
	return &UserResponse{
		ID:        ths.UserID,
		Email:     ths.Email,
		FirstName: ths.FirstName,
		LastName:  ths.LastName,
		Profile: ProfileResponse{
			ID:         ths.ProfileID,
			Job:        ths.Job,
			Activities: ths.Activities,
			Photo:      ths.ProfilePhoto,
		},
	}
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
	Profile      ProfileCreateRequest `json:"profile,omitempty"`
}

type ProfileCreateRequest struct {
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

type TokenPairResponse struct {
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
