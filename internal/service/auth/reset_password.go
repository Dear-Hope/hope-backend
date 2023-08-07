package auth

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/pkg/mailer"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

func (s *service) ResetPassword(ctx context.Context, email string) *response.ServiceError {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &response.ServiceError{
				Code: http.StatusNotFound,
				Msg:  constant.ErrorUserNotFound,
				Err:  err,
			}
		}
		return &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetUserFailed,
			Err:  err,
		}
	}

	user.Password = ""
	key, err := constructKey(*user)
	if err != nil {
		return &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorInternalServer,
			Err:  err,
		}
	}

	if err := s.mailer.Send(ctx, mailer.EmailTemplate{
		Subject:    "Reset Password",
		To:         email,
		From:       "no-reply@dearhope.id",
		SenderName: "Dear Hope",
		Body: fmt.Sprintf(
			`<h4>Halo, %s!</h4>
			</br>
			<p>Seseorang baru-baru ini meminta pengaturan ulang kata sandi untuk akun Dear Hope Anda. Silakan gunakan tautan ini untuk mengatur ulang kata sandi Anda:</p>
			</br>
			<a href="https://dearhope.id/reset?key=%s"> <img alt="reset button" src="https://res.cloudinary.com/shirotama/image/upload/v1656645972/image/static/reset_button.png" width=200/> </a>
			</br>
			<p>*Note: Jika Anda tidak meminta pengaturan ulang kata sandi, Anda dapat mengabaikan email ini.</p>`,
			user.Name,
			key,
		),
	}); err != nil {
		return &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorInternalServer,
			Err:  err,
		}
	}

	return nil
}
