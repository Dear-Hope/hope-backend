package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ReportPost(c echo.Context) error {
	var req model.ReportRequest
	var res model.Response

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)

	}

	report, svcErr := ths.svc.ReportPost(req, c.Get("userID").(uint))
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = report
	return c.JSON(http.StatusOK, res)
}