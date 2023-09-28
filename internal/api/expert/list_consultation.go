package expert

import (
	"HOPE-backend/internal/entity/consultation"
	"HOPE-backend/internal/entity/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) ListConsultation(c echo.Context) error {
	var (
		res    response.Response
		svcErr *response.ServiceError
	)

	res.Result, svcErr = h.consulSvc.GetByExpert(c.Request().Context(), consultation.ExpertListRequest{
		UserId:       0,
		ExpertId:     c.Get("id").(uint64),
		BookingDate:  c.QueryParam("date"),
		BookingMonth: c.QueryParam("month"),
		Status:       consultation.GetStatus(c.QueryParam("status")),
	})
	if svcErr != nil {
		c.Logger().Errorf("[ExpertHandler.ListConsultation]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}
