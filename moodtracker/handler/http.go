package handler

import (
	authMiddleware "HOPE-backend/auth/handler/middleware"
	"HOPE-backend/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type handler struct {
	svc models.MoodTrackerService
}

func NewMoodTrackerHandler(router *gin.RouterGroup, svc models.MoodTrackerService) {
	handler := &handler{
		svc: svc,
	}

	mood := router.Group("mood")
	{
		mood.POST("/", authMiddleware.AuthorizeTokenJWT, handler.NewMoodTrack)
		mood.GET("/", authMiddleware.AuthorizeTokenJWT, handler.ListMoodTrack)
	}
}

func (ths *handler) NewMoodTrack(c *gin.Context) {
	var req models.NewEmotionRequest
	var res models.Response

	currentUserID := c.GetUint("userID")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	emotion, err := ths.svc.NewEmotion(req, currentUserID)
	if err != nil {
		res.Error = err.Error()
		if strings.Contains(err.Error(), "already exists") ||
			strings.Contains(err.Error(), "not found") ||
			strings.Contains(err.Error(), "emotion was not a") {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.JSON(http.StatusInternalServerError, res)
		}
		return
	}

	res.Result = emotion
	c.JSON(http.StatusOK, res)
}

func (ths *handler) ListMoodTrack(c *gin.Context) {
	var res models.Response

	currentUserID := c.GetUint("userID")

	emotions, err := ths.svc.ListEmotion(currentUserID)
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusNotFound, res)
		return
	}

	res.Result = emotions
	c.JSON(http.StatusOK, res)
}
