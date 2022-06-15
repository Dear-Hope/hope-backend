package handler

import (
	"HOPE-backend/v2/models"
	"HOPE-backend/v2/services/auth/handler/middleware"
	"HOPE-backend/v2/services/auth/helper"
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
		auth.POST("/activate", handler.ActivateAccount)
	}
	user := router.Group("/user")
	{
		user.GET("/me", middleware.AuthorizeTokenJWT, handler.GetUserMe)
		user.PUT("/me", middleware.AuthorizeTokenJWT, handler.UpdateUserMe)
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
		} else if res.Error == "account has not been activated yet" {
			c.JSON(http.StatusUnauthorized, res)
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
			userID := rtClaims["userID"].(float64)
			profileID := rtClaims["profileID"].(float64)
			isActive := rtClaims["isActive"].(bool)
			newTokenPair, err := helper.GenerateTokenPair(uint(userID), uint(profileID), isActive)
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

func (ths *handler) GetUserMe(c *gin.Context) {
	var res models.Response
	userID := c.GetUint("userID")

	user, err := ths.svc.GetLoggedInUser(userID)
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusNotFound, res)
		return
	}

	res.Result = user
	c.JSON(http.StatusOK, res)
}

func (ths *handler) UpdateUserMe(c *gin.Context) {
	var res models.Response
	var req models.UpdateRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	req.UserID = c.GetUint("userID")
	req.ProfileID = c.GetUint("profileID")

	updatedUser, err := ths.svc.UpdateLoggedInUser(req)
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res.Result = updatedUser
	c.JSON(http.StatusOK, res)
}

func (ths *handler) ActivateAccount(c *gin.Context) {
	var res models.Response
	var req models.ActivateRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := ths.svc.Activate(req)
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res.Result = token
	c.JSON(http.StatusOK, res)
}
