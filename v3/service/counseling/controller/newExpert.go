package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ths *controller) NewExpert(c echo.Context) error {
	var req model.CreateUpdateExpertRequest
	var res model.Response

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	emotion, svcErr := ths.svc.CreateExpert(req)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = emotion
	return c.JSON(http.StatusOK, res)
}
