package jwt

import (
	"HOPE-backend/config"
	"HOPE-backend/internal/entity/auth"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenClaim struct {
	Id         uint64
	Role       string
	IsVerified bool
}

func GenerateTokenPair(req TokenClaim) (*auth.TokenPairResponse, error) {
	access := jwt.New(jwt.SigningMethodHS256)

	atClaims := access.Claims.(jwt.MapClaims)
	atClaims["access"] = true
	atClaims["id"] = req.Id
	atClaims["isVerified"] = req.IsVerified
	atClaims["role"] = req.Role
	atClaims["expires"] = time.Now().Add(config.Get().Jwt.AccessExpiryInHour * time.Hour).Unix()

	at, err := access.SignedString([]byte(config.Get().Server.SecretKey))
	if err != nil {
		return nil, fmt.Errorf("error generate access token: %v", err)
	}

	refresh := jwt.New(jwt.SigningMethodHS256)

	rtClaims := refresh.Claims.(jwt.MapClaims)
	rtClaims["refresh"] = true
	rtClaims["id"] = req.Id
	rtClaims["isVerified"] = req.IsVerified
	rtClaims["role"] = req.Role
	rtClaims["expires"] = time.Now().Add(config.Get().Jwt.RefreshExpiryInHour * time.Hour).Unix()

	rt, err := refresh.SignedString([]byte(config.Get().Server.SecretKey))
	if err != nil {
		return nil, fmt.Errorf("error generate refresh token: %v", err)
	}

	return &auth.TokenPairResponse{
		Access:  at,
		Refresh: rt,
	}, nil
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
			Id:         uint64(claims["id"].(float64)),
			Role:       claims["role"].(string),
			IsVerified: claims["isVerified"].(bool),
		}, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
