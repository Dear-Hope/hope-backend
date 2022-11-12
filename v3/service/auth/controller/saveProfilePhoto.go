package controller

import (
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func (ths *controller) UploadProfilePhoto(c echo.Context) error {
	var res model.Response

	header, err := c.FormFile("profile_photo")
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	file, err := header.Open()
	if err != nil {
		res.Error = fmt.Sprintf("invalid file: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	userID, _ := c.Get("userID").(uint)

	req := model.SaveProfilePhotoRequest{
		File:      &file,
		Extension: filepath.Ext(header.Filename),
		UserID:    userID,
	}

	link, svcErr := ths.svc.SaveProfilePhoto(req)
	if svcErr != nil {
		res.Error = svcErr.Err.Error()
		return c.JSON(svcErr.Code, res)

	}

	res.Result = map[string]string{
		"link": link,
	}
	return c.JSON(http.StatusOK, res)
}
