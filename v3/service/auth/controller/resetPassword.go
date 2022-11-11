package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ResetPassword(c echo.Context) error {
	var res model.Response
	var req model.ResetPasswordRequest

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	svcErr := ths.svc.ResetPassword(req)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Result = map[string]bool{"success": true}
	return c.JSON(http.StatusOK, res)
}
