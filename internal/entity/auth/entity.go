package auth

// Requests Section
type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Source   string `json:"source"`
	}

	ChangePasswordRequest struct {
		Key         string `json:"key"`
		NewPassword string `json:"new_password"`
	}
)

// Responses Section
type (
	// TokenPairResponse is response struct for access & refresh Token
	TokenPairResponse struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	}
)
