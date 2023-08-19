package expert

import (
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/schedule"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) UpdateSchedule(c echo.Context) error {
	var (
		res response.Response
		req schedule.UpdateRequest
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	req.ExpertId = c.Get("id").(uint64)

	success, svcErr := h.scheduleSvc.Update(c.Request().Context(), req)
	if svcErr != nil {
		c.Logger().Errorf("[ExpertHandler.UpdateSchedule]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	res.Result = map[string]bool{"success": success}
	return c.JSON(http.StatusOK, res)
}
