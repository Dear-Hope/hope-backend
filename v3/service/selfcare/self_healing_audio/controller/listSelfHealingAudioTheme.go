package controller

import (
	"HOPE-backend/v3/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListSelfHealingAudioTheme(c echo.Context) error {
	var res model.Response

	selfHealingAudioThemes, svcErr := ths.svc.ListTheme()
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = selfHealingAudioThemes
	return c.JSON(http.StatusOK, res)
}
