package controller

import (
	"HOPE-backend/v3/middleware"
	"HOPE-backend/v3/service/selfcare/breathing_exercise"

	"github.com/labstack/echo/v4"
)

type controller struct {
	svc breathing_exercise.Service
}

func NewController(router *echo.Group, svc breathing_exercise.Service) {
	controller := &controller{
		svc: svc,
	}

	breathingExercise := router.Group("/selfcare/breathing-exercise")
	{
		breathingExercise.GET("", controller.ListBreathingExercises, middleware.AuthorizeTokenJWT)
		breathingExercise.POST("/last-played", controller.SetLastPlayed, middleware.AuthorizeTokenJWT)
		breathingExercise.GET("/last-played", controller.GetLastPlayed, middleware.AuthorizeTokenJWT)
	}
}
