package controller

import (
	"HOPE-backend/v3/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) GetUserMe(c echo.Context) error {
	var res model.Response
	userID := c.Get("userID")

	user, svcErr := ths.svc.GetLoggedInUser(userID.(uint))
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = user
	return c.JSON(http.StatusOK, res)
}
