package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ChangePassword(c echo.Context) error {
	var res model.Response
	var req model.ChangePasswordRequest

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	token, svcErr := ths.svc.ChangePassword(req)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Result = token
	return c.JSON(http.StatusOK, res)
}
