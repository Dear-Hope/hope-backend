package auth

import (
	"HOPE-backend/v2/models"
	"HOPE-backend/v2/services/auth/helper"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/pquerna/otp/totp"
	sendblue "github.com/sendinblue/APIv3-go-library/lib"
)

type service struct {
	repo   models.AuthRepository
	mailer *sendblue.APIClient
	cache  *cache.Cache
}

func NewAuthService(repo models.AuthRepository, mailer *sendblue.APIClient, cache *cache.Cache) models.AuthService {
	return &service{
		repo:   repo,
		mailer: mailer,
		cache:  cache,
	}
}

func (ths *service) Login(req models.LoginRequest) (*models.TokenPair, error) {
	user, err := ths.repo.GetUserWithProfileByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = helper.ComparePassword([]byte(req.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	if !user.IsActive {
		return nil, errors.New("account has not been activated yet")
	}

	tokenPair, err := helper.GenerateTokenPair(user.UserID, user.ProfileID, user.IsActive)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func generateSecretKey(email string) (string, error) {
	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      "DearHope",
			AccountName: email,
		},
	)
	if err != nil {
		return "", err
	}

	return key.Secret(), nil
}

func (ths *service) Register(req models.RegisterRequest) (*models.TokenPair, error) {
	hashedPassword, err := helper.EncryptPassword([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	secret, err := generateSecretKey(req.Email)
	if err != nil {
		return nil, err
	}

	newUser := models.DBUserWithProfile{
		Email:        req.Email,
		Password:     hashedPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		ProfilePhoto: req.ProfilePhoto,
		Job:          req.Profile.Job,
		Activities:   req.Profile.Activities,
		SecretKey:    secret,
	}
	userID, profileID, err := ths.repo.CreateUserWithProfile(newUser)
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(userID, profileID, false)
	if err != nil {
		return nil, err
	}

	newUser.UserID = userID
	newUser.ProfileID = profileID
	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return nil, err
	}

	template := models.EmailTemplate{
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
			newUser.FirstName+" "+newUser.LastName,
			code,
		),
	}

	err = sendKey(ths.mailer, template)
	if err != nil {
		return nil, err
	}

	ths.cache.SetDefault(req.Email, 0)

	return tokenPair, nil
}

func constructKey(user models.DBUserWithProfile) (string, error) {
	userByte, err := json.Marshal(user)
	if err != nil {
		return "", errors.New("failed to construct key: " + err.Error())
	}

	key, err := helper.Encrypt(string(userByte))
	if err != nil {
		return "", errors.New("failed to encrypt key: " + err.Error())
	}

	return key, nil
}

func sendKey(mailer *sendblue.APIClient, template models.EmailTemplate) error {
	_, _, err := mailer.TransactionalEmailsApi.SendTransacEmail(
		context.Background(),
		sendblue.SendSmtpEmail{
			Sender: &sendblue.SendSmtpEmailSender{
				Name:  "Dear Hope",
				Email: "no-reply@dearhope.id",
			},
			To: []sendblue.SendSmtpEmailTo{
				{
					Email: template.Email,
				},
			},
			Subject:     template.Subject,
			HtmlContent: template.Content,
		},
	)
	if err != nil {
		return errors.New("failed to send activation key: " + err.Error())
	}

	return nil
}

func (ths *service) GetLoggedInUser(id uint) (*models.UserResponse, error) {
	user, err := ths.repo.GetUserWithProfileByID(id)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		User: models.User{
			ID:        user.UserID,
			Email:     user.Email,
			Password:  "",
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Profile: models.Profile{
			ID:         user.ProfileID,
			Job:        user.Job,
			Activities: user.Activities,
			Photo:      user.ProfilePhoto,
		},
	}, nil
}

func (ths *service) UpdateLoggedInUser(req models.UpdateRequest) (*models.UserResponse, error) {
	newPassword, err := helper.EncryptPassword([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	newUser := models.DBUserWithProfile{
		UserID:       req.UserID,
		ProfileID:    req.ProfileID,
		Email:        req.Email,
		Password:     newPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		ProfilePhoto: req.ProfilePhoto,
		Job:          req.Profile.Job,
		Activities:   req.Profile.Activities,
	}

	updatedUser, err := ths.repo.UpdateUserWithProfile(newUser)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		User: models.User{
			ID:        updatedUser.UserID,
			Email:     updatedUser.Email,
			Password:  "",
			FirstName: updatedUser.FirstName,
			LastName:  updatedUser.LastName,
		},
		Profile: models.Profile{
			ID:         updatedUser.ProfileID,
			Job:        updatedUser.Job,
			Activities: updatedUser.Activities,
			Photo:      updatedUser.ProfilePhoto,
		},
	}, nil
}

func (ths *service) Activate(req models.ActivateRequest) (*models.TokenPair, error) {
	user, err := ths.repo.GetUserWithProfileByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !totp.Validate(req.Code, user.SecretKey) {
		return nil, errors.New("your activation code has expired")
	}

	err = ths.repo.SetUserToActive(user.UserID)
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(user.UserID, user.ProfileID, true)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (ths *service) ResetPassword(req models.ResetPasswordRequest) error {
	user, err := ths.repo.GetUserWithProfileByEmail(req.Email)
	if err != nil {
		return err
	}

	user.Password = ""
	key, err := constructKey(*user)
	if err != nil {
		return err
	}

	template := models.EmailTemplate{
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
		return err
	}

	return nil
}

func (ths *service) ChangePassword(req models.ChangePasswordRequest) (*models.TokenPair, error) {
	userString, err := helper.Decrypt(strings.TrimPrefix(req.Key, "https://dearhope.id/reset?key="))
	if err != nil {
		return nil, errors.New("failed to decrypt change password key: " + err.Error())
	}

	var user models.DBUserWithProfile
	err = json.Unmarshal([]byte(userString), &user)
	if err != nil {
		return nil, errors.New("failed to change password: " + err.Error())
	}

	hashedPassword, err := helper.EncryptPassword([]byte(req.NewPassword))
	if err != nil {
		return nil, err
	}

	err = ths.repo.UpdatePassword(user.UserID, hashedPassword)
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(user.UserID, user.ProfileID, true)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (ths *service) SaveProfilePhoto(req models.SaveProfilePhotoRequest) (string, error) {
	dirPath := fmt.Sprintf("assets/users/%d/profile_photo%s", req.UserID, req.Extension)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll("assets/users/"+fmt.Sprint(req.UserID), 0770)
		if err != nil {
			return "", fmt.Errorf("failed to save profile photo: %s", err.Error())
		}
	}

	out, err := os.Create(dirPath)
	if err != nil {
		return "", fmt.Errorf("failed to save profile photo: %s", err.Error())
	}

	defer out.Close()
	_, err = io.Copy(out, *req.File)
	if err != nil {
		return "", fmt.Errorf("failed to save profile photo: %s", err.Error())
	}

	filepath := "https://13.251.114.0.nip.io/" + dirPath
	err = ths.repo.SetUserProfilePhoto(req.UserID, filepath)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func (ths *service) ResendActivationCode(req models.ResetPasswordRequest) error {
	user, err := ths.repo.GetUserWithProfileByEmail(req.Email)
	if err != nil {
		return err
	}

	if user.IsActive {
		return errors.New("user already activated the account")
	}

	if retries, found := ths.cache.Get(req.Email); found {
		if retries.(int) >= 3 {
			return errors.New("resend limit reached, please try again after a while")
		}
		if err := ths.cache.Increment(req.Email, 1); err != nil {
			return err
		}
	}

	code, err := totp.GenerateCode(user.SecretKey, time.Now())
	if err != nil {
		return err
	}

	template := models.EmailTemplate{
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
		return err
	}

	return nil
}

func (ths *service) DeleteUser(req models.ResetPasswordRequest) error {
	err := ths.repo.DeleteUserByEmail(req.Email)
	if err != nil {
		return err
	}

	return nil
}
