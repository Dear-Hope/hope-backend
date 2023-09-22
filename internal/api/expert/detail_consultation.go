package expert

import (
	"HOPE-backend/internal/entity/response"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) DetailConsultation(c echo.Context) error {
	var (
		res    response.Response
		svcErr *response.ServiceError
	)

	consulId, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Result, svcErr = h.consulSvc.GetDetailByExpert(c.Request().Context(), consulId)
	if svcErr != nil {
		c.Logger().Errorf("[ExpertHandler.DetailConsultation]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}
