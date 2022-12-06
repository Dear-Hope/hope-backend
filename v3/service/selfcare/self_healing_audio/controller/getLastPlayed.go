package controller

import (
	"HOPE-backend/v3/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) GetLastPlayed(c echo.Context) error {
	var res model.Response

	breathingExercise, svcErr := ths.svc.GetLastPlayed(c.Get("userID").(uint))
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = breathingExercise
	return c.JSON(http.StatusOK, res)
}
