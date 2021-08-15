package handler

import (
	authMiddleware "HOPE-backend/auth/handler/middleware"
	"HOPE-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	svc models.AmbulanceService
}

func NewAmbulanceHandler(router *gin.RouterGroup, svc models.AmbulanceService) {
	handler := &handler{
		svc: svc,
	}

	ambulance := router.Group("ambulance")
	{
		ambulance.GET("/", authMiddleware.AuthorizeTokenJWT, handler.GetAmbulanceList)
	}

}

func (ths *handler) GetAmbulanceList(c *gin.Context) {
	var res models.Response
	search := c.Query("search")
	location := c.Query("location")

	ambulances, err := ths.svc.List(search, location)
	if err != nil {
		res.Error = err.Error()
	}

	res.Result = ambulances
	c.JSON(http.StatusOK, res)
}
