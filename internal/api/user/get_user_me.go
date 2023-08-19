package user

import (
	"HOPE-backend/internal/entity/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetUserMe(c echo.Context) error {
	var (
		res    response.Response
		svcErr *response.ServiceError
	)
	userId := c.Get("id")

	res.Result, svcErr = h.svc.Get(c.Request().Context(), userId.(uint64))
	if svcErr != nil {
		c.Logger().Errorf("[UserHandler.GetUserMe]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}
