package auth

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

	VerifyRequest struct {
		Email string `json:"email"`
		Code  string `json:"code"`
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
