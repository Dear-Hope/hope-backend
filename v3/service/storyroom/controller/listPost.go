package controller

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/storyroom/filter"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListPost(c echo.Context) error {
	var res model.Response

	categoryID, err := getListPostFilters(c)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	posts, svcErr := ths.svc.List(
		filter.List{
			CategoryID: categoryID,
			Sort:       c.QueryParam("sort"),
			Direction:  c.QueryParam("direction"),
		},
		c.Get("userID").(uint),
	)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = posts
	return c.JSON(http.StatusOK, res)
}

func getListPostFilters(c echo.Context) (uint, error) {
	var (
		_categoryID uint64
		err         error
	)

	if c.QueryParam("categoryID") != "" {
		_categoryID, err = strconv.ParseUint(c.QueryParam("categoryID"), 10, 0)
		if err != nil {
			return 0, err
		}
	}

	categoryID := uint(_categoryID)

	return categoryID, nil
}
