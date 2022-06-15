package handler

import (
	authMiddleware "HOPE-backend/v1/auth/handler/middleware"
	"HOPE-backend/v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	svc models.LaboratoryService
}

func NewLaboratoryHandler(router *gin.RouterGroup, svc models.LaboratoryService) {
	handler := &handler{
		svc: svc,
	}

	laboratory := router.Group("laboratory")
	{
		laboratory.GET("/", authMiddleware.AuthorizeTokenJWT, handler.GetLaboratoryList)
	}

}

func (ths *handler) GetLaboratoryList(c *gin.Context) {
	res := gin.H{}

	search := c.Query("search")
	location := c.Query("location")
	filter := c.Query("filter")

	laboratories, err := ths.svc.List(search, location, filter)
	if err != nil {
		res["error"] = err.Error()
	}

	res["result"] = laboratories
	c.JSON(http.StatusOK, res)
}
