package consultation

import (
	"HOPE-backend/internal/entity/consultation"
	"HOPE-backend/internal/entity/response"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateStatus(c echo.Context) error {
	var (
		req    consultation.UpdateStatusRequest
		res    response.Response
		svcErr *response.ServiceError
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	consulId, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Result, svcErr = h.svc.UpdateStatus(c.Request().Context(), consulId, consultation.GetStatus(req.Status))
	if svcErr != nil {
		c.Logger().Errorf("[ExpertHandler.DetailConsultation]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}
