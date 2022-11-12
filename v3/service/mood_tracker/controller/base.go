package controller

import (
	"HOPE-backend/v3/middleware"
	"HOPE-backend/v3/service/mood_tracker"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controller struct {
	svc mood_tracker.Service
}

func NewController(router *echo.Group, svc mood_tracker.Service) {
	controller := &controller{
		svc: svc,
	}

	mood := router.Group("/mood")
	{
		mood.GET("", controller.ListMoodData)
		mood.POST("", controller.NewMood, middleware.AuthorizeTokenJWT)
		mood.GET("/today", controller.ListMoodToday, middleware.AuthorizeTokenJWT)
		mood.GET("/week", controller.ListMoodWeekly, middleware.AuthorizeTokenJWT)
		mood.GET("/month", controller.ListMoodMonthly, middleware.AuthorizeTokenJWT)
	}
}

func getParams(c echo.Context) (int, int, error) {
	var (
		month, year int
		err         error
	)

	if c.QueryParam("month") != "" {
		month, err = strconv.Atoi(c.QueryParam("month"))
		if err != nil {
			return 0, 0, err
		}
	}

	if c.QueryParam("year") != "" {
		year, err = strconv.Atoi(c.QueryParam("year"))
		if err != nil {
			return 0, 0, err
		}
	}

	return month, year, nil
}
