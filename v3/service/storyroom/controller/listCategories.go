package controller

import (
	"HOPE-backend/v3/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListCategories(c echo.Context) error {
	var res model.Response

	categories, svcErr := ths.svc.ListCategories()
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = categories
	return c.JSON(http.StatusOK, res)
}
