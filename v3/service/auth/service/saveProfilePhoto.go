package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (ths *service) SaveProfilePhoto(req model.SaveProfilePhotoRequest) (string, *model.ServiceError) {
	dirPath := fmt.Sprintf("assets/users/%d/profile_photo%s", req.UserID, req.Extension)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll("assets/users/"+fmt.Sprint(req.UserID), 0770)
		if err != nil {
			return "", &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  fmt.Errorf("%s: %s", constant.ERROR_SAVE_PROFILE_PHOTO_FAILED, err.Error()),
			}
		}
	}

	out, err := os.Create(dirPath)
	if err != nil {
		return "", &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("%s: %s", constant.ERROR_SAVE_PROFILE_PHOTO_FAILED, err.Error()),
		}
	}

	defer out.Close()
	_, err = io.Copy(out, *req.File)
	if err != nil {
		return "", &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("%s: %s", constant.ERROR_SAVE_PROFILE_PHOTO_FAILED, err.Error()),
		}
	}

	filepath := "https://13.251.114.0.nip.io/" + dirPath
	err = ths.repo.SetProfilePhoto(req.UserID, filepath)
	if err != nil {
		return "", &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_SAVE_PROFILE_PHOTO_FAILED),
		}
	}

	return filepath, nil
}
