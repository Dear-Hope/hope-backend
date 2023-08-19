package user

import (
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/user"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) Verify(c echo.Context) error {
	var (
		res    response.Response
		req    user.VerifyRequest
		svcErr *response.ServiceError
	)
	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Result, svcErr = h.svc.Verify(c.Request().Context(), req)
	if svcErr != nil {
		c.Logger().Errorf("[AuthHandler.VerifyAccount]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}
