package controller

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/counseling/filter"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListExperts(c echo.Context) error {
	var res model.Response

	topicId, err := getListExpertsFilter(c)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	experts, svcErr := ths.svc.ListExperts(
		filter.ListExpert{
			TopicId:   topicId,
			Expertise: c.QueryParam("expertise"),
		},
	)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = experts
	return c.JSON(http.StatusOK, res)
}

func getListExpertsFilter(c echo.Context) (uint, error) {
	var (
		_topicId uint64
		err      error
	)

	if c.QueryParam("topicId") != "" {
		_topicId, err = strconv.ParseUint(c.QueryParam("topicId"), 10, 0)
		if err != nil {
			return 0, err
		}
	}

	topicId := uint(_topicId)

	return topicId, nil
}
