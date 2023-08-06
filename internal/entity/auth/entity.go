package auth

// User is representation of table users, user_profiles, and user_roles
type User struct {
	Id               uint64 `db:"id"`
	Email            string `db:"email"`
	Password         string `db:"password"`
	Name             string `db:"name"`
	Alias            string `db:"alias"`
	IsVerified       bool   `db:"is_verified"`
	SecretKey        string `db:"secret_key"`
	Photo            string `db:"photo"`
	TotalAudioPlayed int    `db:"total_audio_played"`
	TotalTimePlayed  int    `db:"total_time_played"`
	LongestStreak    int    `db:"longest_streak"`
	Role             string `db:"role"`
}

// Requests Section
type (
	// RegisterRequest is request struct for registering user
	RegisterRequest struct {
		Email        string `json:"email"`
		Password     string `json:"password"`
		Role         string `json:"role"`
		Name         string `json:"name,omitempty"`
		Alias        string `json:"alias,omitempty"`
		ProfilePhoto string `json:"profilePhoto,omitempty"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

// Responses Section
type (
	// UserResponse is response struct for User model
	UserResponse struct {
		ID               uint64 `json:"id" `
		Email            string `json:"email"`
		Name             string `json:"name"`
		Alias            string `json:"alias"`
		Photo            string `json:"photo"`
		Role             string `json:"role"`
		TotalAudioPlayed int    `json:"totalAudioPlayed"`
		TotalTimePlayed  int    `json:"totalTimePlayed"`
		LongestStreak    int    `json:"longest_streak"`
		IsVerified       bool   `json:"isVerified"`
	}

	// TokenPairResponse is response struct for access & refresh Token
	TokenPairResponse struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	}
)
