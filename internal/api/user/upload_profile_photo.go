package user

import (
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/user"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"path/filepath"
)

func (h *Handler) UploadProfilePhoto(c echo.Context) error {
	var res response.Response

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

	link, svcErr := h.svc.SaveProfilePhoto(c.Request().Context(), user.SaveProfilePhotoRequest{
		Id:        c.Get("id").(uint64),
		File:      &file,
		Extension: filepath.Ext(header.Filename),
	})
	if svcErr != nil {
		c.Logger().Errorf("[UserHandler.UploadProfilePhoto]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	res.Result = map[string]string{"link": link}
	return c.JSON(http.StatusOK, res)
}
