package controller

import (
	"HOPE-backend/v3/middleware"
	"HOPE-backend/v3/service/selfcare/category"

	"github.com/labstack/echo/v4"
)

type controller struct {
	svc category.Service
}

func NewController(router *echo.Group, svc category.Service) {
	controller := &controller{
		svc: svc,
	}

	category := router.Group("/selfcare/category")
	{
		category.GET("", controller.ListCategory, middleware.AuthorizeTokenJWT)
	}
}
