package user

import (
	"mime/multipart"
	"time"
)

// User is representation of table users, user_profiles, and user_roles
type User struct {
	Id               uint64    `db:"id"`
	Email            string    `db:"email"`
	Password         string    `db:"password"`
	Name             string    `db:"name"`
	Alias            string    `db:"alias"`
	IsVerified       bool      `db:"is_verified"`
	SecretKey        string    `db:"secret_key"`
	Photo            string    `db:"photo"`
	TotalAudioPlayed int       `db:"total_audio_played"`
	TotalTimePlayed  int       `db:"total_time_played"`
	LongestStreak    int       `db:"longest_streak"`
	Role             string    `db:"role_name"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
	IsDeleted        bool      `db:"is_deleted"`
}

// Requests section
type (
	// RegisterRequest is request struct for registering user
	RegisterRequest struct {
		Email        string `json:"email"`
		Password     string `json:"password"`
		Name         string `json:"name,omitempty"`
		Alias        string `json:"alias,omitempty"`
		ProfilePhoto string `json:"profilePhoto,omitempty"`
	}

	UpdateRequest struct {
		Id           uint64 `json:"id"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		Name         string `json:"name"`
		Alias        string `json:"alias"`
		ProfilePhoto string `json:"profilePhoto"`
	}

	VerifyRequest struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	SaveProfilePhotoRequest struct {
		Id        uint64
		File      *multipart.File
		Extension string
	}
)

// Responses Section
type (
	// Response is response struct for User model
	Response struct {
		Id               uint64 `json:"id" `
		Email            string `json:"email"`
		Name             string `json:"name"`
		Alias            string `json:"alias"`
		Photo            string `json:"photo"`
		TotalAudioPlayed int    `json:"totalAudioPlayed,omitempty"`
		TotalTimePlayed  int    `json:"totalTimePlayed,omitempty"`
		LongestStreak    int    `json:"longest_streak,omitempty"`
		IsVerified       bool   `json:"isVerified,omitempty"`
	}
)

// Functions Section

// ToResponse is helper function to generate User -> Response
func (u User) ToResponse() *Response {
	return &Response{
		Id:               u.Id,
		Email:            u.Email,
		Name:             u.Name,
		Alias:            u.Alias,
		Photo:            u.Photo,
		TotalAudioPlayed: u.TotalAudioPlayed,
		TotalTimePlayed:  u.TotalTimePlayed,
		LongestStreak:    u.LongestStreak,
	}
}
