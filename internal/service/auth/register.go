package auth

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/response"
	"context"
	"net/http"
	"strings"
)

func (ths *service) Register(ctx context.Context, req auth.RegisterRequest) (*auth.TokenPairResponse, *response.ServiceError) {
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
	user, err := ths.repo.CreateUser(ctx, newUser)
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

	//code, err := totp.GenerateCode(secret, time.Now())
	//if err != nil {
	//	return nil, &response.ServiceError{
	//		Code: http.StatusInternalServerError,
	//		Err:  errors.New(constant.ERROR_GENERATE_OTP_CODE),
	//	}
	//}
	//
	//template := response.EmailTemplate{
	//	Subject: "Account Activation",
	//	Email:   req.Email,
	//	Content: fmt.Sprintf(
	//		`<h4>Hai, %s!</h4>
	//		</br>
	//		<p>Selamat datang di keluarga Dear Hope. Mulai detik ini kamu tidak sendiri lagi, karena ada Hope yang menemani. Sebelum kita memulai, kita hanya butuh untuk mengkonfirmasi bahwa ini adalah kamu, silahkan masukkan kode OTP di bawah ini:</p>
	//		</br>
	//		<h3>%s</h3>
	//		</br>
	//		<p>Semoga harimu menyenangkan</p>`,
	//		user.Name,
	//		code,
	//	),
	//}
	//
	//err = sendKey(ths.mailer, template)
	//if err != nil {
	//	return nil, &response.ServiceError{
	//		Code: http.StatusInternalServerError,
	//		Err:  err,
	//	}
	//}
	//
	//ths.cache.SetDefault(req.Email, 0)

	return tokenPair, nil
}
