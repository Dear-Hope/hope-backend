package user

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/user"
	"HOPE-backend/pkg/helpers"
	"context"
	"net/http"
)

func (s *service) Update(ctx context.Context, req user.UpdateRequest) (bool, *response.ServiceError) {
	var (
		newPassword string
		err         error
	)
	if req.Password != "" {
		newPassword, err = helpers.EncryptPassword([]byte(req.Password))
		if err != nil {
			return false, &response.ServiceError{
				Code: http.StatusInternalServerError,
				Msg:  constant.ErrorUpdateUserFailed,
				Err:  err,
			}
		}
	}

	newUser := user.User{
		Id:       req.Id,
		Email:    req.Email,
		Password: newPassword,
		Name:     req.Name,
		Alias:    req.Alias,
		Photo:    req.ProfilePhoto,
	}

	_, err = s.repo.UpdateUser(ctx, newUser)
	if err != nil {
		return false, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorUpdateUserFailed,
			Err:  err,
		}
	}

	return true, nil
}
