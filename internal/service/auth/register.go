package auth

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/pkg/mailer"
	"context"
	"fmt"
	"github.com/pquerna/otp/totp"
	"net/http"
	"strings"
	"time"
)

func (s *service) Register(ctx context.Context, req auth.RegisterRequest) (*auth.TokenPairResponse, *response.ServiceError) {
	hashedPassword, err := encryptPassword([]byte(req.Password))
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorInternalServer,
			Err:  err,
		}
	}

	secret, err := generateSecretKey(req.Email)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorInternalServer,
			Err:  err,
		}
	}

	newUser := auth.User{
		Email:      req.Email,
		Password:   hashedPassword,
		Name:       req.Name,
		Alias:      req.Alias,
		IsVerified: req.Role == "EXPERT",
		SecretKey:  secret,
		Photo:      req.ProfilePhoto,
		Role:       req.Role,
	}
	user, err := s.repo.CreateUser(ctx, newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &response.ServiceError{
				Code: http.StatusBadRequest,
				Msg:  constant.ErrorUserAlreadyExists,
				Err:  err,
			}
		} else {
			return nil, &response.ServiceError{
				Code: http.StatusInternalServerError,
				Msg:  constant.ErrorCreateUserFailed,
				Err:  err,
			}
		}
	}

	tokenPair, err := generateTokenPair(user.Id, user.Role, user.IsVerified)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGenerateToken,
			Err:  err,
		}
	}

	if !user.IsVerified {
		code, err := totp.GenerateCode(secret, time.Now())
		if err != nil {
			return nil, &response.ServiceError{
				Code: http.StatusInternalServerError,
				Msg:  constant.ErrorGenerateOtpCode,
				Err:  fmt.Errorf("[AuthSvc][010008] failed to generate otp code: %v", err),
			}
		}

		err = s.mailer.Send(ctx, mailer.EmailTemplate{
			Subject:    "Account Activation",
			To:         req.Email,
			From:       "no-reply@dearhope.id",
			SenderName: "Dear Hope",
			Body: fmt.Sprintf(
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
		})
		if err != nil {
			return nil, &response.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  err,
			}
		}

		s.cache.Set(ctx, req.Email, 0, 0)
	}

	return tokenPair, nil
}
