package main

import (
	"HOPE-backend/auth"
	_authHandler "HOPE-backend/auth/handler"
	_authRepo "HOPE-backend/auth/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := NewPostgreSQLDatabase()

	v1 := router.Group("/api/v1")

	authRepo := _authRepo.NewPostgreSQLRepository(db)
	authSvc := auth.NewAuthService(authRepo)
	_authHandler.NewAuthHandler(v1, authSvc)

	router.Run(":80")
}
