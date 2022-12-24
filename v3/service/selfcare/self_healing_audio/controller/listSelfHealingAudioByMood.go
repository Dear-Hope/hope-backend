package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListSelfHealingAudioByMood(c echo.Context) error {
	var (
		res    model.Response
		moodID uint64
		err    error
	)

	if c.QueryParam("moodID") != "" {
		moodID, err = strconv.ParseUint(c.QueryParam("moodID"), 10, 0)
		if err != nil {
			res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
			return c.JSON(http.StatusBadRequest, res)
		}
	}

	selfHealingAudioTheme, svcErr := ths.svc.ListAudioByMood(uint(moodID), c.Get("userID").(uint))
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = selfHealingAudioTheme
	return c.JSON(http.StatusOK, res)
}
