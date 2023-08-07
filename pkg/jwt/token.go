package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenClaim struct {
	UserId     uint64
	Role       string
	IsVerified bool
}

func ValidateToken(encodedToken, secretKey string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}
		return []byte(secretKey), nil
	})
}

func AuthorizeToken(tokenString, secretKey string, isRefresh bool) (*TokenClaim, error) {
	token, err := ValidateToken(tokenString, secretKey)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if isRefresh && !claims["refresh"].(bool) {
			return nil, errors.New("token was not a refresh token")
		}

		if int64(claims["expires"].(float64)) < time.Now().Unix() {
			return nil, errors.New("token has expired")
		}
		if !claims["isVerified"].(bool) {
			return nil, errors.New("account has not been activated yet")
		}

		return &TokenClaim{
			UserId:     uint64(claims["userId"].(float64)),
			Role:       claims["role"].(string),
			IsVerified: claims["isVerified"].(bool),
		}, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
