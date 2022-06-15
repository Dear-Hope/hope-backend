package middleware

import (
	"HOPE-backend/v2/models"
	"HOPE-backend/v2/services/auth/helper"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeTokenJWT(c *gin.Context) {
	var res models.Response

	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		res.Error = "authorization header not given"
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := helper.ValidateToken(tokenString)
	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if !claims["isActive"].(bool) {
			res.Error = "account has not been activated yet"
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		}
		userID := claims["userID"].(float64)
		profileID := claims["profileID"].(float64)
		c.Set("userID", uint(userID))
		c.Set("profileID", uint(profileID))
		c.Next()
	} else {
		res.Error = "invalid token"
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
	}
}
