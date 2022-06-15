package handler

import (
	"HOPE-backend/v2/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type handler struct {
	svc models.NewsletterService
}

func NewNewsletterService(router *gin.RouterGroup, svc models.NewsletterService) {
	handler := &handler{
		svc: svc,
	}

	newsletter := router.Group("newsletter")
	{
		newsletter.POST("/", handler.NewSubscription)
		newsletter.DELETE("/:email", handler.DeleteSubscription)
	}
}

func (ths *handler) NewSubscription(c *gin.Context) {
	var req models.NewSubscriberRequest
	var res models.Response

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = ths.svc.Subscribe(req)
	if err != nil {
		res.Error = err.Error()
		if strings.Contains(err.Error(), "invalid") {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.JSON(http.StatusInternalServerError, res)
		}
		return
	}

	res.Result = map[string]bool{"success": true}
	c.JSON(http.StatusOK, res)
}

func (ths *handler) DeleteSubscription(c *gin.Context) {
	var res models.Response

	email := c.Param("email")

	err := ths.svc.Unsubscribe(email)
	if err != nil {
		res.Error = err.Error()
		if strings.Contains(err.Error(), "invalid") {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.JSON(http.StatusInternalServerError, res)
		}
		return
	}

	res.Result = map[string]bool{"success": true}
	c.JSON(http.StatusOK, res)
}
