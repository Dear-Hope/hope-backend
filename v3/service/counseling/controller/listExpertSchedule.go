package controller

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/counseling/filter"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (ths *controller) ListExpertSchedule(c echo.Context) error {
	var res model.Response

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

	offset, err := strconv.ParseInt(c.QueryParam("offset"), 10, 0)
	if err != nil {
		res.Error = "mohon sertakan zona waktu"
		return c.JSON(http.StatusBadRequest, res)
	}

	if c.QueryParam("date") == "" {
		res.Error = "mohon sertakan tanggal konsultasi"
		return c.JSON(http.StatusBadRequest, res)
	}

	consuls, svcErr := ths.svc.ListExpertSchedule(
		filter.ListExpertSchedule{
			ExpertId: expertId,
			TypeId:   typeId,
			Date:     c.QueryParam("date"),
			Offset:   int(offset),
		},
	)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = consuls
	return c.JSON(http.StatusOK, res)
}
