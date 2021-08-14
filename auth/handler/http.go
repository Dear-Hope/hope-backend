package handler

import (
	"HOPE-backend/auth/handler/helper"
	"HOPE-backend/models"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type handler struct {
	svc models.AuthService
}

func NewAuthHandler(router *gin.RouterGroup, svc models.AuthService) {
	handler := &handler{
		svc: svc,
	}

	auth := router.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)
		auth.POST("/login/refresh", handler.RefreshToken)

	}
}

func (ths *handler) Login(c *gin.Context) {
	var req models.LoginRequest
	var res models.Response

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := ths.svc.Login(req)
	if err != nil {
		res.Error = err.Error()
		if res.Error == "failed to generate token" {
			c.JSON(http.StatusInternalServerError, res)
		} else {
			c.JSON(http.StatusNotFound, res)
		}
		return
	}

	res.Result = token
	c.JSON(http.StatusOK, res)
}

func (ths *handler) Register(c *gin.Context) {
	var req models.RegisterRequest
	var res models.Response

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := ths.svc.Register(req)
	if err != nil {
		res.Error = err.Error()
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.JSON(http.StatusInternalServerError, res)
		}
		return
	}

	res.Result = token
	c.JSON(http.StatusCreated, res)
}

func (ths *handler) RefreshToken(c *gin.Context) {
	var tokenPair models.TokenPair
	var res models.Response

	c.ShouldBindJSON(&tokenPair)
	refreshToken, err := helper.ValidateToken(tokenPair.Refresh)
	if err != nil {
		res.Error = err.Error()
		log.Printf("error validate refresh token: %s", err.Error())

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if rtClaims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		if rtClaims["refresh"] == true {
			newTokenPair, err := helper.GenerateTokenPair(uint(rtClaims["id"].(float64)))
			if err != nil {
				res.Error = err.Error()
				log.Printf("error generate new access token: %s", err.Error())
				c.JSON(http.StatusInternalServerError, res)
				return
			}
			newTokenPair.Refresh = ""
			res.Result = newTokenPair
			c.JSON(http.StatusOK, res)
			return
		}
		res.Error = "token was not a refresh token"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res.Error = "invalid refresh token"
	c.JSON(http.StatusUnauthorized, res)
}
