package controller

import (
	"HOPE-backend/v3/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListReason(c echo.Context) error {
	var res model.Response

	reasons, svcErr := ths.svc.ListReason()
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = reasons
	return c.JSON(http.StatusOK, res)
}