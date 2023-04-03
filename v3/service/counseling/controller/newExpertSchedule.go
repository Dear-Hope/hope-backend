package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (ths *controller) NewExpertSchedule(c echo.Context) error {
	var req model.NewConsultationRequest
	var res model.Response

	expertId, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	err = c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	rowsAffected, svcErr := ths.svc.CreateExpertSchedule(req, expertId)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = map[string]int64{"totalCreated": rowsAffected}
	return c.JSON(http.StatusOK, res)
}
