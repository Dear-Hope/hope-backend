package controller

import (
	"HOPE-backend/v3/middleware"
	"HOPE-backend/v3/service/selfcare/self_healing_audio"

	"github.com/labstack/echo/v4"
)

type controller struct {
	svc self_healing_audio.Service
}

func NewController(router *echo.Group, svc self_healing_audio.Service) {
	controller := &controller{
		svc: svc,
	}

	selfHealingAudio := router.Group("/selfcare/audio")
	{
		selfHealingAudio.GET("/:id", controller.GetSelfHealingAudio, middleware.AuthorizeTokenJWT)
		selfHealingAudio.GET("/theme", controller.ListSelfHealingAudioTheme, middleware.AuthorizeTokenJWT)
		selfHealingAudio.GET("/theme/:id", controller.GetSelfHealingAudioTheme, middleware.AuthorizeTokenJWT)
		selfHealingAudio.POST("/last-played", controller.SetLastPlayed, middleware.AuthorizeTokenJWT)
		selfHealingAudio.GET("/last-played", controller.GetLastPlayed, middleware.AuthorizeTokenJWT)
	}
}
