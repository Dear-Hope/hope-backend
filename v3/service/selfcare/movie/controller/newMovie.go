package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) NewMovie(c echo.Context) error {
	var req model.MovieRequest
	var res model.Response

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)

	}

	movie, svcErr := ths.svc.Create(req)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = movie
	return c.JSON(http.StatusOK, res)
}
