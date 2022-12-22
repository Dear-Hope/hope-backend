package controller

import (
	"HOPE-backend/v3/middleware"
	"HOPE-backend/v3/service/storyroom"

	"github.com/labstack/echo/v4"
)

type controller struct {
	svc storyroom.Service
}

func NewController(router *echo.Group, svc storyroom.Service) {
	controller := &controller{
		svc: svc,
	}

	storyroom := router.Group("/story-room")
	{
		storyroom.POST("/post", controller.NewPost, middleware.AuthorizeTokenJWT)
		storyroom.GET("/post", controller.ListPost, middleware.AuthorizeTokenJWT)
		storyroom.GET("/post/me", controller.ListMyPost, middleware.AuthorizeTokenJWT)
		storyroom.GET("/post/:id", controller.DetailPost, middleware.AuthorizeTokenJWT)
		storyroom.DELETE("/post/:id", controller.DeletePost, middleware.AuthorizeTokenJWT)
		storyroom.POST("/post/:id/comment", controller.NewComment, middleware.AuthorizeTokenJWT)
		storyroom.GET("/post/:id/like", controller.LikePost, middleware.AuthorizeTokenJWT)
		storyroom.GET("/categories", controller.ListCategories, middleware.AuthorizeTokenJWT)
	}
}
