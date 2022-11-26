package controller

import (
	"HOPE-backend/v3/service/newsletter"

	"github.com/labstack/echo/v4"
)

type controller struct {
	svc newsletter.Service
}

func NewController(router *echo.Group, svc newsletter.Service) {
	controller := &controller{
		svc: svc,
	}

	newsletter := router.Group("newsletter")
	{
		newsletter.POST("", controller.NewSubscription)
		newsletter.DELETE("/:email", controller.DeleteSubscription)
	}
}
