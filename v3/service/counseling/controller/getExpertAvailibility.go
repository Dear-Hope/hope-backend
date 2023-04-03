package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (ths *controller) GetExpertAvailability(c echo.Context) error {
	var (
		res model.Response
	)
	expertId, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	isAvailable, svcErr := ths.svc.GetExpertAvailability(uint(expertId))
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = map[string]bool{"isAvailable": isAvailable}
	return c.JSON(http.StatusOK, res)
}
