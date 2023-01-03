package controller

import (
	"HOPE-backend/v3/middleware"
	"HOPE-backend/v3/service/auth"

	"github.com/labstack/echo/v4"
)

type controller struct {
	svc auth.Service
}

func NewController(router *echo.Group, svc auth.Service) {
	controller := &controller{
		svc: svc,
	}

	auth := router.Group("/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
		auth.POST("/login/refresh", controller.RefreshToken)
		auth.POST("/activate", controller.ActivateAccount)
		auth.POST("/resend", controller.ResendActivationCode)
		auth.POST("/password/reset", controller.ResetPassword)
		auth.POST("/password/change", controller.ChangePassword)
	}
	user := router.Group("/user")
	{
		user.GET("/me", controller.GetUserMe, middleware.AuthorizeTokenJWT)
		user.PUT("/me", controller.UpdateUserMe, middleware.AuthorizeTokenJWT)
		user.POST("/me/upload/photo", controller.UploadProfilePhoto, middleware.AuthorizeTokenJWT)
		user.POST("/delete", controller.DeleteUser)
		user.POST("/block", controller.BlockUser, middleware.AuthorizeTokenJWT)
	}
}
