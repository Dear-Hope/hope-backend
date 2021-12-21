package handler

import (
	authMiddleware "HOPE-backend/auth/handler/middleware"
	"HOPE-backend/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type handler struct {
	svc      models.ChatService
	upgrader websocket.Upgrader
}

func NewChatHandler(router *gin.RouterGroup, svc models.ChatService, upgrader websocket.Upgrader) {
	handler := &handler{
		svc:      svc,
		upgrader: upgrader,
	}

	chat := router.Group("conversation")
	{
		chat.POST("/", authMiddleware.AuthorizeTokenJWT, handler.StartConversation)
		chat.GET("/", authMiddleware.AuthorizeTokenJWT, handler.ListConversation)
		chat.GET("/:id", handler.GetConversation)
		chat.POST("/:id/chat", authMiddleware.AuthorizeTokenJWT, handler.SendChat)
		chat.GET("/ws", handler.ServeChatWS)
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

func (ths *handler) GetConversation(c *gin.Context) {
	var res models.Response
	strID := c.Param("id")
	currentUserID := c.GetUint("userID")

	conversationID, err := strconv.Atoi(strID)
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	conversation, err := ths.svc.GetConversation(uint(conversationID))
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusNotFound, res)
		return
	}

	if currentUserID != conversation.FirstUserID && currentUserID != conversation.SecondUserID {
		res.Error = "You are not the owner of this conversation"
		c.JSON(http.StatusForbidden, res)
		return
	}

	res.Result = conversation
	c.JSON(http.StatusOK, res)
}

func (ths *handler) ListConversation(c *gin.Context) {
	var res models.Response
	strID := c.Query("userID")
	currentUserID := c.GetUint("userID")

	userID, err := strconv.Atoi(strID)
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if currentUserID != uint(userID) {
		res.Error = "You are not the owner of these conversations"
		c.JSON(http.StatusForbidden, res)
		return
	}

	conversations, err := ths.svc.ListConversation(uint(userID))
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusNotFound, res)
		return
	}

	res.Result = conversations
	c.JSON(http.StatusOK, res)
}

func (ths *handler) SendChat(c *gin.Context) {
	var req models.NewChatRequest
	var res models.Response

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	chat, err := ths.svc.NewChat(req)
	if err != nil {
		res.Error = err.Error()
		if strings.Contains(err.Error(), "already exists") ||
			strings.Contains(err.Error(), "not found") ||
			strings.Contains(err.Error(), "conversation owners") {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.JSON(http.StatusInternalServerError, res)
		}
		return
	}

	res.Result = chat
	c.JSON(http.StatusOK, res)
}
