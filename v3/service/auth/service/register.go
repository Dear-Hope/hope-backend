package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/auth/helper"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
)

func (ths *service) Register(req model.RegisterRequest) (*model.TokenPairResponse, *model.ServiceError) {
	hashedPassword, err := helper.EncryptPassword([]byte(req.Password))
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	secret, err := generateSecretKey(req.Email)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	newUser := model.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		SecretKey: secret,
		Profile: model.Profile{
			ProfilePhoto: req.ProfilePhoto,
			Job:          req.Profile.Job,
			Activities:   req.Profile.Activities,
		},
	}
	user, err := ths.repo.Create(newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &model.ServiceError{
				Code: http.StatusBadRequest,
				Err:  errors.New(constant.ERROR_USER_ALREADY_EXISTS),
			}
		} else {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  errors.New(constant.ERROR_CREATE_USER_FAILED),
			}
		}
	}

	tokenPair, err := helper.GenerateTokenPair(user.UserID, user.ProfileID, false)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GENERATE_TOKEN),
		}
	}

	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return nil, &model.ServiceError{
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
			user.FirstName+" "+user.LastName,
			code,
		),
	}

	err = sendKey(ths.mailer, template)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	ths.cache.SetDefault(req.Email, 0)

	return tokenPair, nil
}
