package handler

import (
	"HOPE-backend/v2/models"
	authMiddleware "HOPE-backend/v2/services/auth/handler/middleware"
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

	selfcare := router.Group("selfcare")
	{
		selfcare.POST("/", authMiddleware.AuthorizeTokenJWT, handler.NewSelfCareItem)
		selfcare.GET("/", authMiddleware.AuthorizeTokenJWT, handler.ListSelfCareItems)
		selfcare.GET("/types", authMiddleware.AuthorizeTokenJWT, handler.ListSelfCareTypes)
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

func (ths *handler) ListSelfCareItems(c *gin.Context) {
	var res models.Response
	var items []*models.SelfCareItem
	var err error

	mood := c.Query("mood")

	if mood == "" {
		items, err = ths.svc.ListItems()
	} else {
		items, err = ths.svc.GetItemsByMood(mood)
	}
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusNotFound, res)
		return
	}

	res.Result = items
	c.JSON(http.StatusOK, res)
}

func (ths *handler) ListSelfCareTypes(c *gin.Context) {
	var res models.Response
	var err error

	types, err := ths.svc.ListTypes()
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusNotFound, res)
		return
	}

	res.Result = types
	c.JSON(http.StatusOK, res)
}
