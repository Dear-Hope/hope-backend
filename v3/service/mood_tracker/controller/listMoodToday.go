package controller

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/mood_tracker/filter"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListMoodToday(c echo.Context) error {
	var res model.Response

	currentUserID := c.Get("userID").(uint)

	periodFrom, periodTo := getToday()

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

func getToday() (int64, int64) { //get first day of the month
	t := time.Now()
	first := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	last := first.AddDate(0, 0, 1).Add(time.Nanosecond * -1)

	return first.UnixMilli(), last.UnixMilli()
}
