package controller

import (
	"HOPE-backend/v3/middleware"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/auth/helper"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func (ths *controller) RefreshToken(c echo.Context) error {
	var tokenPair model.TokenPairResponse
	var res model.Response

	c.Bind(&tokenPair)
	refreshToken, err := middleware.ValidateToken(tokenPair.Refresh)
	if err != nil {
		res.Error = err.Error()
		log.Printf("error validate refresh token: %s", err.Error())

		return c.JSON(http.StatusBadRequest, res)
	}

	if rtClaims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		if rtClaims["refresh"] == true {
			if int64(rtClaims["expires"].(float64)) < time.Now().Unix() {
				res.Error = "refresh token has expired"
				return c.JSON(http.StatusUnauthorized, res)
			}
			userID := rtClaims["userID"].(float64)
			profileID := rtClaims["profileID"].(float64)
			isActive := rtClaims["isActive"].(bool)
			newTokenPair, err := helper.GenerateTokenPair(uint(userID), uint(profileID), isActive)
			if err != nil {
				res.Error = err.Error()
				log.Printf("error generate new access token: %s", err.Error())
				return c.JSON(http.StatusInternalServerError, res)
			}
			newTokenPair.Refresh = ""
			res.Result = newTokenPair
			return c.JSON(http.StatusOK, res)
		}
		res.Error = "token was not a refresh token"
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Error = "invalid refresh token"
	return c.JSON(http.StatusUnauthorized, res)
}
