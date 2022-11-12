package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pquerna/otp/totp"
)

func (ths *service) ResendActivationCode(req model.ResetPasswordRequest) *model.ServiceError {
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

	if user.IsActive {
		return &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ERROR_ACCOUNT_ALREADY_ACTIVATED),
		}
	}

	if retries, found := ths.cache.Get(req.Email); found {
		if retries.(int) >= 3 {
			return &model.ServiceError{
				Code: http.StatusBadRequest,
				Err:  errors.New(constant.ERROR_RESEND_LIMIT_REACHED),
			}
		}
		if err := ths.cache.Increment(req.Email, 1); err != nil {
			return &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  err,
			}
		}
	}

	code, err := totp.GenerateCode(user.SecretKey, time.Now())
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GENERATE_OTP_CODE),
		}
	}

	template := model.EmailTemplate{
		Subject: "Account Activation",
		Email:   req.Email,
		Content: fmt.Sprintf(
			`<h4>Hai, %s!</h4>
			</br>
			<p>Selamat datang di keluarga Dear Hope. Mulai detik ini kamu tidak sendiri lagi, karena ada Hope yang menemani. Sebelum kita memulai, kita hanya butuh untuk mengkonfirmasi bahwa ini adalah kamu, silahkan masukkan kode OTP di bawah ini:</p>
			</br>
			<h3>%s</h3>
			</br>
			<p>Semoga harimu menyenangkan</p>`,
			user.Name,
			code,
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
