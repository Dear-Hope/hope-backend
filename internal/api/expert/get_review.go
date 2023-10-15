package expert

import (
	"HOPE-backend/internal/entity/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetReview(c echo.Context) error {
	var (
		res    response.Response
		svcErr *response.ServiceError
	)

	res.Result, svcErr = h.reviewSvc.Get(c.Request().Context(), c.Get("id").(uint64))
	if svcErr != nil {
		c.Logger().Errorf("[ExpertHandler.GetReview]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}
