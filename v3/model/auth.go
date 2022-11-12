package model

import "mime/multipart"

type (
	User struct {
		UserID    uint   `db:"user_id"`
		Email     string `db:"email"`
		Password  string `db:"password"`
		Name      string `db:"name"`
		Alias     string `db:"alias"`
		IsActive  bool   `db:"is_active"`
		SecretKey string `db:"secret_key"`

		Profile
	}

	Profile struct {
		ProfileID        uint   `db:"profile_id"`
		ProfilePhoto     string `db:"photo"`
		Job              string `db:"job"`
		Activities       string `db:"activities"`
		TotalAudioPlayed int    `db:"total_audio_played"`
		TotalTimePlayed  int    `db:"total_time_played"`
		LongestStreak    int    `db:"longest_streak"`
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
		ID      uint            `json:"id" `
		Email   string          `json:"email"`
		Name    string          `json:"name"`
		Alias   string          `json:"alias"`
		Profile ProfileResponse `json:"profile,omitempty"`
	}

	ProfileResponse struct {
		ID               uint   `json:"id"`
		Job              string `json:"job"`
		Activities       string `json:"activities"`
		Photo            string `json:"photo"`
		TotalAudioPlayed int    `json:"total_audio_played"`
		TotalTimePlayed  int    `json:"total_time_played"`
		LongestStreak    int    `json:"longest_streak"`
	}
)

func (ths User) ToUserResponse() *UserResponse {
	return &UserResponse{
		ID:    ths.UserID,
		Email: ths.Email,
		Name:  ths.Name,
		Alias: ths.Alias,
		Profile: ProfileResponse{
			ID:               ths.ProfileID,
			Job:              ths.Job,
			Activities:       ths.Activities,
			Photo:            ths.ProfilePhoto,
			TotalAudioPlayed: ths.TotalAudioPlayed,
			TotalTimePlayed:  ths.TotalTimePlayed,
			LongestStreak:    ths.LongestStreak,
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
	Name         string               `json:"name,omitempty"`
	Alias        string               `json:"alias,omitempty"`
	ProfilePhoto string               `json:"profile_photo,omitempty"`
	Profile      ProfileCreateRequest `json:"profile,omitempty"`
}

type ProfileCreateRequest struct {
	Job              string `json:"job,omitempty"`
	Activities       string `json:"activities,omitempty"`
	TotalAudioPlayed int    `json:"total_audio_played,omitempty"`
	TotalTimePlayed  int    `json:"total_time_played,omitempty"`
	LongestStreak    int    `json:"longest_streak,omitempty"`
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
