package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ths *controller) GetSelfHealingAudioTheme(c echo.Context) error {
	var res model.Response

	themeID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	selfHealingAudioTheme, svcErr := ths.svc.GetTheme(uint(themeID), c.Get("userID").(uint))
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = selfHealingAudioTheme
	return c.JSON(http.StatusOK, res)
}
