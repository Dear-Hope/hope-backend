package controller

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListMovie(c echo.Context) error {
	var res model.Response

	moodID, needID, err := getListMovieFilters(c)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	movies, svcErr := ths.svc.List(filter.ListMovie{MoodID: moodID, NeedID: needID})
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = movies
	return c.JSON(http.StatusOK, res)
}

func getListMovieFilters(c echo.Context) (*uint, *uint, error) {
	var (
		_moodID, _needID uint64
		err              error
	)

	if c.QueryParam("moodID") != "" {
		_moodID, err = strconv.ParseUint(c.QueryParam("moodID"), 10, 0)
		if err != nil {
			return nil, nil, err
		}
	}

	if c.QueryParam("needID") != "" {
		_needID, err = strconv.ParseUint(c.QueryParam("needID"), 10, 0)
		if err != nil {
			return nil, nil, err
		}
	}

	moodID := uint(_moodID)
	needID := uint(_needID)

	return &moodID, &needID, nil
}
