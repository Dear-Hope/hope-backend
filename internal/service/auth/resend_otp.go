package auth

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/pkg/mailer"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/pquerna/otp/totp"
	"net/http"
	"time"
)

func (s *service) ResendOtp(ctx context.Context, email string) *response.ServiceError {
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

	if user.IsVerified {
		return &response.ServiceError{
			Code: http.StatusBadRequest,
			Msg:  constant.ErrorAccountAlreadyVerified,
			Err:  fmt.Errorf("[AuthSvc.ResendOtp][010014] account already verified"),
		}
	}

	if retries, found := s.cache.Get(ctx, email); found {
		if retries.(int) >= 3 {
			return &response.ServiceError{
				Code: http.StatusBadRequest,
				Msg:  constant.ErrorResendLimitReached,
				Err:  fmt.Errorf("[AuthSvc.ResendOtp][010015] resend limit reached"),
			}
		}
		if err := s.cache.Increment(ctx, email, 1); err != nil {
			return &response.ServiceError{
				Code: http.StatusInternalServerError,
				Msg:  constant.ErrorInternalServer,
				Err:  err,
			}
		}
	}

	code, err := totp.GenerateCode(user.SecretKey, time.Now())
	if err != nil {
		return &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGenerateOtpCode,
			Err:  fmt.Errorf("[AuthSvc.ResendOtp][010016] failed to generate otp code: %v", err),
		}
	}

	err = s.mailer.Send(ctx, mailer.EmailTemplate{
		Subject:    "Account Activation",
		To:         email,
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
		return &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorInternalServer,
			Err:  err,
		}
	}

	return nil
}
