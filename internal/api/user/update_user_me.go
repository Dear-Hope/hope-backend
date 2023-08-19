package user

import (
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/user"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) UpdateUserMe(c echo.Context) error {
	var (
		res response.Response
		req user.UpdateRequest
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	req.Id = c.Get("id").(uint64)

	success, svcErr := h.svc.Update(c.Request().Context(), req)
	if svcErr != nil {
		c.Logger().Errorf("[UserHandler.UpdateUserMe]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	res.Result = map[string]bool{"success": success}
	return c.JSON(http.StatusOK, res)
}
