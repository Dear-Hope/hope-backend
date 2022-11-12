package controller

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/mood_tracker/filter"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListMoodMonthly(c echo.Context) error {
	var res model.Response

	currentUserID := c.Get("userID").(uint)

	month, year, err := getParams(c)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	periodFrom, periodTo := getFirstAndLastDayOfMonth(time.Month(month), year)

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

func getFirstAndLastDayOfMonth(month time.Month, year int) (int64, int64) { //get first day of the month
	if month == 0 && year == 0 {
		t := time.Now()
		month = t.Month()
		year = t.Year()
	}

	firstday := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	lastday := firstday.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

	return firstday.UnixMilli(), lastday.UnixMilli()
}
