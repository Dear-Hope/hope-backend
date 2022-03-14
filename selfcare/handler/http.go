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
	svc models.SelfCareService
}

func NewSelfCareHandler(router *gin.RouterGroup, svc models.SelfCareService) {
	handler := &handler{
		svc: svc,
	}

	mood := router.Group("selfcare")
	{
		mood.POST("/", authMiddleware.AuthorizeTokenJWT, handler.NewSelfCareItem)
		mood.GET("/", authMiddleware.AuthorizeTokenJWT, handler.ListSelfCareItemsByMood)
	}
}

func (ths *handler) NewSelfCareItem(c *gin.Context) {
	var req models.NewSelfCareItemRequest
	var res models.Response

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	selfCareItem, err := ths.svc.NewItem(req)
	if err != nil {
		res.Error = err.Error()
		if strings.Contains(err.Error(), "already exists") ||
			strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.JSON(http.StatusInternalServerError, res)
		}
		return
	}

	res.Result = selfCareItem
	c.JSON(http.StatusOK, res)
}

func (ths *handler) ListSelfCareItemsByMood(c *gin.Context) {
	var res models.Response

	mood := c.Query("mood")

	items, err := ths.svc.GetItemsByMood(mood)
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusNotFound, res)
		return
	}

	res.Result = items
	c.JSON(http.StatusOK, res)
}
