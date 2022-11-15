package middleware

import (
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AuthorizeTokenJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var res model.Response

		const BEARER_SCHEMA = "Bearer "
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			res.Error = "authorization header not given"
			return c.JSON(http.StatusUnauthorized, res)
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := ValidateToken(tokenString)
		if err != nil {
			res.Error = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if int64(claims["expires"].(float64)) < time.Now().Unix() {
				res.Error = "access token has expired"
				return c.JSON(http.StatusUnauthorized, res)
			}
			if !claims["isActive"].(bool) {
				res.Error = "account has not been activated yet"
				return c.JSON(http.StatusUnauthorized, res)
			}
			userID := claims["userID"].(float64)
			profileID := claims["profileID"].(float64)
			c.Set("userID", uint(userID))
			c.Set("profileID", uint(profileID))
			return next(c)
		} else {
			res.Error = "invalid token"
			return c.JSON(http.StatusUnauthorized, res)
		}
	}
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}
		return []byte("hope-secret-key"), nil
	})
}
