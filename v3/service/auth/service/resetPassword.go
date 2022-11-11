package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

func (ths *service) ResetPassword(req model.ResetPasswordRequest) *model.ServiceError {
	user, err := ths.repo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &model.ServiceError{
				Code: http.StatusNotFound,
				Err:  errors.New(constant.ERROR_USER_NOT_FOUND),
			}
		}
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_USER_FAILED),
		}
	}

	user.Password = ""
	key, err := constructKey(*user)
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	template := model.EmailTemplate{
		Subject: "Reset Password",
		Email:   req.Email,
		Content: fmt.Sprintf(
			`<h4>Halo, %s!</h4>
			</br>
			<p>Seseorang baru-baru ini meminta pengaturan ulang kata sandi untuk akun Dear Hope Anda. Silakan gunakan tautan ini untuk mengatur ulang kata sandi Anda:</p>
			</br>
			<a href="https://dearhope.id/reset?key=%s"> <img alt="reset button" src="https://res.cloudinary.com/shirotama/image/upload/v1656645972/image/static/reset_button.png" width=200/> </a>
			</br>
			<p>*Note: Jika Anda tidak meminta pengaturan ulang kata sandi, Anda dapat mengabaikan email ini.</p>`,
			user.FirstName+" "+user.LastName,
			key,
		),
	}

	err = sendKey(ths.mailer, template)
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return nil
}
