package expert

import (
	"HOPE-backend/internal/entity/response"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) GetScheduleUser(c echo.Context) error {
	var (
		res    response.Response
		svcErr *response.ServiceError
	)

	expertId, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	typeId, err := strconv.ParseUint(c.QueryParam("typeId"), 10, 0)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	date, err := strconv.ParseInt(c.QueryParam("date"), 10, 0)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Result, svcErr = h.scheduleSvc.GetTimeslotUsers(c.Request().Context(), expertId, typeId, date)
	if svcErr != nil {
		c.Logger().Errorf("[ExpertHandler.GetScheduleUser]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}
