package user

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/user"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (s *service) SaveProfilePhoto(ctx context.Context, req user.SaveProfilePhotoRequest) (
	string, *response.ServiceError) {
	dirPath := fmt.Sprintf("assets/users/%d/profile_photo%s", req.Id, req.Extension)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll("assets/users/"+fmt.Sprint(req.Id), 0770)
		if err != nil {
			return "", &response.ServiceError{
				Code: http.StatusInternalServerError,
				Msg:  constant.ErrorSaveProfilePhotoFailed,
				Err:  fmt.Errorf("[AuthSvc.SaveProfilePhoto][010023] %s", err.Error()),
			}
		}
	}

	out, err := os.Create(dirPath)
	if err != nil {
		return "", &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorSaveProfilePhotoFailed,
			Err:  fmt.Errorf("[AuthSvc.SaveProfilePhoto][010024] %s", err.Error()),
		}
	}
	defer func() {
		_ = out.Close()
	}()

	_, err = io.Copy(out, *req.File)
	if err != nil {
		return "", &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorSaveProfilePhotoFailed,
			Err:  fmt.Errorf("[AuthSvc.SaveProfilePhoto][010025] %s", err.Error()),
		}
	}

	filepath := "https://108.136.238.205.nip.io/" + dirPath
	_, err = s.repo.UpdateUser(ctx, user.User{
		Id:    req.Id,
		Photo: dirPath,
	})
	if err != nil {
		return "", &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorSaveProfilePhotoFailed,
			Err:  err,
		}
	}

	return filepath, nil
}
