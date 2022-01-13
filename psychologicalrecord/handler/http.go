package handler

import (
	authMiddleware "HOPE-backend/auth/handler/middleware"
	"HOPE-backend/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type handler struct {
	svc models.PsychologicalRecordService
}

func NewPsychologicalRecordHandler(router *gin.RouterGroup, svc models.PsychologicalRecordService) {
	handler := &handler{
		svc: svc,
	}

	chat := router.Group("record")
	{
		chat.POST("/", authMiddleware.AuthorizeTokenJWT, handler.NewRecord)
		chat.GET("/", authMiddleware.AuthorizeTokenJWT, handler.ListRecord)
		chat.GET("/:id", authMiddleware.AuthorizeTokenJWT, handler.GetRecord)
	}
}

func (ths *handler) NewRecord(c *gin.Context) {
	var req models.PsychologicalRecord
	var res models.Response

	currentUserID := c.GetUint("userID")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if currentUserID != req.PsychologistID {
		res.Error = "cannot create new record for other psychologist"
		c.JSON(http.StatusForbidden, res)
		return
	}

	if currentUserID == req.PatientID {
		res.Error = "patient cannot create a new psychological record"
		c.JSON(http.StatusForbidden, res)
		return
	}

	record, err := ths.svc.NewRecord(req)
	if err != nil {
		res.Error = err.Error()
		if strings.Contains(err.Error(), "already exists") ||
			strings.Contains(err.Error(), "not found") ||
			strings.Contains(err.Error(), "record was not a") {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.JSON(http.StatusInternalServerError, res)
		}
		return
	}

	res.Result = record
	c.JSON(http.StatusOK, res)
}

func (ths *handler) GetRecord(c *gin.Context) {
	var res models.Response

	currentUserID := c.GetUint("userID")
	strID := c.Param("id")

	recordID, err := strconv.Atoi(strID)
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	record, err := ths.svc.GetRecord(uint(recordID))
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusNotFound, res)
		return
	}

	if currentUserID != record.PsychologistID {
		res.Error = "you are not the owner of this psychological record"
		c.JSON(http.StatusForbidden, res)
		return
	}

	if currentUserID == record.PatientID {
		res.Error = "patient cannot see pyschological record"
		c.JSON(http.StatusForbidden, res)
		return
	}

	res.Result = record
	c.JSON(http.StatusOK, res)
}

func (ths *handler) ListRecord(c *gin.Context) {
	var res models.Response

	currentUserID := c.GetUint("userID")

	records, err := ths.svc.ListRecord(currentUserID)
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusNotFound, res)
		return
	}

	res.Result = records
	c.JSON(http.StatusOK, res)
}
