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
	svc models.ChatService
}

func NewChatHandler(router *gin.RouterGroup, svc models.ChatService) {
	handler := &handler{
		svc: svc,
	}

	chat := router.Group("conversation")
	{
		chat.POST("/", authMiddleware.AuthorizeTokenJWT, handler.StartConversation)
	}

}

func (ths *handler) StartConversation(c *gin.Context) {
	var req models.NewConversationRequest
	var res models.Response

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	conversation, err := ths.svc.NewConversation(req)
	if err != nil {
		res.Error = err.Error()
		if strings.Contains(err.Error(), "already exists") || strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.JSON(http.StatusInternalServerError, res)
		}
		return
	}

	res.Result = conversation
	c.JSON(http.StatusCreated, res)
}
