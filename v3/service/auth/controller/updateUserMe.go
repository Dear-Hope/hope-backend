package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ths *controller) UpdateUserMe(c echo.Context) error {
	var res model.Response
	var req model.UpdateRequest

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	userID, _ := c.Get("userID").(uint)
	profileID, _ := c.Get("profileID").(uint)

	req.UserID = userID
	req.ProfileID = profileID

	updatedUser, svcErr := ths.svc.UpdateLoggedInUser(req)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)
	}

	res.Result = updatedUser
	return c.JSON(http.StatusOK, res)
}
