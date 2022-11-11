package controller

import (
	"HOPE-backend/v3/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListMoodData(c echo.Context) error {
	var res model.Response

	moods, svcErr := ths.svc.ListMoodData()
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = moods
	return c.JSON(http.StatusOK, res)
}
