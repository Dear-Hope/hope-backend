package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) BlockUser(c echo.Context) error {
	var (
		res model.Response
		req model.BlockUserRequest
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	userID := c.Get("userID")

	svcErr := ths.svc.BlockUser(userID.(uint), req)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = map[string]bool{"success": true}
	return c.JSON(http.StatusOK, res)
}
