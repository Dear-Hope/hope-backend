package controller

import (
	"HOPE-backend/v3/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) DeleteSubscription(c echo.Context) error {
	var res model.Response

	email := c.Param("email")

	err := ths.svc.Unsubscribe(email)
	if err != nil {
		res.Error = err.Err.Error()
		return c.JSON(err.Code, res)
	}

	res.Result = map[string]bool{"success": true}
	return c.JSON(http.StatusOK, res)
}
