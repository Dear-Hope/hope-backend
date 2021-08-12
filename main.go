package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello", hello)
	}

	router.Run(":8000")
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "Hello, world!"})
}
