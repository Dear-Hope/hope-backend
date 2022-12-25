package controller

import (
	"HOPE-backend/v3/middleware"
	"HOPE-backend/v3/service/selfcare/music"

	"github.com/labstack/echo/v4"
)

type controller struct {
	svc music.Service
}

func NewController(router *echo.Group, svc music.Service) {
	controller := &controller{
		svc: svc,
	}

	music := router.Group("/selfcare/music")
	{
		music.GET("", controller.ListPlaylist, middleware.AuthorizeTokenJWT)
	}
}
