package controller

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/storyroom/filter"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) ListMyPost(c echo.Context) error {
	var res model.Response

	userID := c.Get("userID").(uint)

	posts, svcErr := ths.svc.List(
		filter.List{UserID: userID},
		userID,
	)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = posts
	return c.JSON(http.StatusOK, res)
}
