package helpers

import (
	"HOPE-backend/v3/constant"
	"time"
)

func GetStartOfDayUTCFromDateWithOffset(date string, offset int) time.Time {
	t, _ := time.Parse(constant.FormatDate, date)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, getTimezoneFromOffset(offset)).UTC()
}

func GetEndOfDayUTCFromDateWithOffset(date string, offset int) time.Time {
	t, _ := time.Parse(constant.FormatDate, date)
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, getTimezoneFromOffset(offset)).UTC()
}

func getTimezoneFromOffset(offset int) *time.Location {
	return time.FixedZone("", offset*constant.HourInSec)
}
