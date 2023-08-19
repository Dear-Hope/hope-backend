package expert

import (
	"HOPE-backend/internal/entity/expert"
	"HOPE-backend/internal/entity/response"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) Register(c echo.Context) error {
	var (
		res    response.Response
		req    expert.CreateUpdateRequest
		svcErr *response.ServiceError
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Result, svcErr = h.svc.Create(c.Request().Context(), req)
	if svcErr != nil {
		c.Logger().Errorf("[ExpertHandler.Register]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}
