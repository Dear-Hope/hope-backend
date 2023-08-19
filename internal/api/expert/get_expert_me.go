package expert

import (
	"HOPE-backend/internal/entity/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetExpertMe(c echo.Context) error {
	var (
		res    response.Response
		svcErr *response.ServiceError
	)

	res.Result, svcErr = h.svc.Get(c.Request().Context(), c.Get("id").(uint64))
	if svcErr != nil {
		c.Logger().Errorf("[ExpertHandler.GetExpertMe]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}
