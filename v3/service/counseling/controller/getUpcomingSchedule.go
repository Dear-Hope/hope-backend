package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (ths *controller) GetUpcomingSchedule(c echo.Context) error {
	var (
		res model.Response
	)
	expertId, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	typeId, err := strconv.ParseInt(c.QueryParam("typeId"), 10, 0)
	if err != nil {
		res.Error = "mohon sertakan jenis konsultasi"
		return c.JSON(http.StatusBadRequest, res)
	}

	consul, svcErr := ths.svc.GetUpcomingSchedule(expertId, typeId)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = consul
	return c.JSON(http.StatusOK, res)
}
