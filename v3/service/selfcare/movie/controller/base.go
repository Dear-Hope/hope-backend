package controller

import (
	"HOPE-backend/v3/middleware"
	"HOPE-backend/v3/service/selfcare/movie"

	"github.com/labstack/echo/v4"
)

type controller struct {
	svc movie.Service
}

func NewController(router *echo.Group, svc movie.Service) {
	controller := &controller{
		svc: svc,
	}

	movie := router.Group("/selfcare/movie")
	{
		movie.POST("", controller.NewMovie)
		movie.GET("", controller.ListMovie, middleware.AuthorizeTokenJWT)
		movie.GET("/:id", controller.DetailMovie, middleware.AuthorizeTokenJWT)
	}
}
