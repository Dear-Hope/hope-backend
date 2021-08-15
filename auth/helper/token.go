package helper

import (
	"HOPE-backend/models"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokenPair(user *models.User) (*models.TokenPair, error) {
	generateFailed := errors.New("failed to generate token")

	access := jwt.New(jwt.SigningMethodHS256)

	atClaims := access.Claims.(jwt.MapClaims)
	atClaims["access"] = true
	atClaims["userID"] = user.ID
	atClaims["profileID"] = user.Profile.ID
	atClaims["expires"] = time.Now().Add(15 * time.Minute).Unix()

	at, err := access.SignedString([]byte("hope-secret-key"))
	if err != nil {
		log.Printf("Error generate access token: %s", err.Error())
		return nil, generateFailed
	}

	refresh := jwt.New(jwt.SigningMethodHS256)

	rtClaims := refresh.Claims.(jwt.MapClaims)
	rtClaims["refresh"] = true
	atClaims["userID"] = user.ID
	atClaims["profileID"] = user.Profile.ID
	rtClaims["expires"] = time.Now().Add(24 * time.Hour).Unix()

	rt, err := refresh.SignedString([]byte("hope-secret-key"))
	if err != nil {
		log.Printf("Error generate refresh token: %s", err.Error())
		return nil, generateFailed
	}

	return &models.TokenPair{
		Access:  at,
		Refresh: rt,
	}, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}
		return []byte("hope-secret-key"), nil
	})
}
