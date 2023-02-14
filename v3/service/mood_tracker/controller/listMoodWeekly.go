package controller

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/mood_tracker/filter"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListMoodWeekly(c echo.Context) error {
	var res model.Response

	currentUserID := c.Get("userID").(uint)

	periodFrom, periodTo := getDayOfWeek()

	emotions, svcErr := ths.svc.List(currentUserID, filter.List{
		PeriodFrom: &periodFrom,
		PeriodTo:   &periodTo,
	})
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = emotions
	return c.JSON(http.StatusOK, res)
}

func getDayOfWeek() (int64, int64) { //get monday 00:00:00
	tm := time.Now()
	weekday := time.Duration(tm.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := tm.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return currentZeroDay.Add(-1 * (weekday - 1) * 24 * time.Hour).UnixMilli(),
		currentZeroDay.Add(1*(7-weekday)*24*time.Hour).Add(24*time.Hour).UnixMilli() - 1000
}
