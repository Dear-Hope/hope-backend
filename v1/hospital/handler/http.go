package handler

import (
	authMiddleware "HOPE-backend/v1/auth/handler/middleware"
	"HOPE-backend/v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	svc models.HospitalService
}

func NewHospitalHandler(router *gin.RouterGroup, svc models.HospitalService) {
	handler := &handler{
		svc: svc,
	}

	hospital := router.Group("hospital")
	{
		hospital.GET("/", authMiddleware.AuthorizeTokenJWT, handler.GetHospitalList)
	}

}

func (ths *handler) GetHospitalList(c *gin.Context) {
	res := gin.H{}

	search := c.Query("search")
	location := c.Query("location")
	filter := c.Query("filter")

	hospitals, err := ths.svc.List(search, location, filter)
	if err != nil {
		res["error"] = err.Error()
	}

	res["result"] = hospitals
	c.JSON(http.StatusOK, res)
}
