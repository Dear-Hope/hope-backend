package controller

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListCategory(c echo.Context) error {
	var res model.Response

	categories, svcErr := ths.svc.List(
		filter.ListCategory{
			ExcludeIDs: []uint{
				uint(constant.CATEGORY_AUDIO_SELF_HEALING_ID),
				uint(constant.CATEGORY_MUSIC_ID),
			},
		},
	)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = categories
	return c.JSON(http.StatusOK, res)
}
