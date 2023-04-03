package controller

import (
	"HOPE-backend/v3/middleware"
	"HOPE-backend/v3/service/counseling"

	"github.com/labstack/echo/v4"
)

type controller struct {
	svc counseling.Service
}

func NewController(router *echo.Group, svc counseling.Service) {
	controller := &controller{
		svc: svc,
	}

	counsel := router.Group("/counseling")
	{
		counsel.GET("/topic", controller.ListTopics, middleware.AuthorizeTokenJWT)
		counsel.POST("/expert", controller.NewExpert, middleware.AuthorizeTokenJWT)
		counsel.GET("/expert", controller.ListExperts, middleware.AuthorizeTokenJWT)
		counsel.GET("/expert/:id", controller.DetailExpert, middleware.AuthorizeTokenJWT)
		counsel.PUT("/expert/:id", controller.UpdateExpert, middleware.AuthorizeTokenJWT)
		counsel.DELETE("/expert/:id", controller.DeleteExpert, middleware.AuthorizeTokenJWT)
		counsel.GET("/expert/:id/available", controller.GetExpertAvailability, middleware.AuthorizeTokenJWT)
		counsel.GET("/expert/:id/schedule", controller.ListExpertSchedule, middleware.AuthorizeTokenJWT)
		counsel.POST("/expert/:id/schedule", controller.NewExpertSchedule, middleware.AuthorizeTokenJWT)
		counsel.GET("/expert/:id/schedule/next", controller.GetUpcomingSchedule, middleware.AuthorizeTokenJWT)
		counsel.POST("/expert/:id/schedule/book", controller.BookConsultation, middleware.AuthorizeTokenJWT)
	}
}
