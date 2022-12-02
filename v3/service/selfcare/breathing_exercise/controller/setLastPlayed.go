package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) SetLastPlayed(c echo.Context) error {
	var (
		req model.BreathingExerciseHistoryRequest
		res model.Response
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	svcErr := ths.svc.SetLastPlayed(c.Get("userID").(uint), req)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = map[string]bool{"success": true}
	return c.JSON(http.StatusOK, res)
}
