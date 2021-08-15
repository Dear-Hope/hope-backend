package handler

import (
	authMiddleware "HOPE-backend/auth/handler/middleware"
	"HOPE-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	svc models.MedicineService
}

func NewMedicineHandler(router *gin.RouterGroup, svc models.MedicineService) {
	handler := &handler{
		svc: svc,
	}

	medicine := router.Group("medicine")
	{
		medicine.GET("/", authMiddleware.AuthorizeTokenJWT, handler.GetMedicineList)
	}

}

func (ths *handler) GetMedicineList(c *gin.Context) {
	var res models.Response
	kind := c.Param("kind")

	medicines, err := ths.svc.List(models.Kind(kind))
	if err != nil {
		res.Error = err.Error()
	}

	res.Result = medicines
	c.JSON(http.StatusOK, res)
}
