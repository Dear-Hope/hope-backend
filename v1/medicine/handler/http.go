package handler

import (
	authMiddleware "HOPE-backend/v1/auth/handler/middleware"
	"HOPE-backend/v1/models"
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
	res := gin.H{}
	kind := c.Query("kind")

	medicines, err := ths.svc.List(models.Kind(kind))
	if err != nil {
		res["error"] = err.Error()
	}

	res["result"] = medicines
	c.JSON(http.StatusOK, res)
}
