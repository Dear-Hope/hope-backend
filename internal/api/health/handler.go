package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct{}

func (h *Handler) Check(c echo.Context) error {
	return c.JSON(http.StatusOK, "Server is OK!")
}
