package auth

import (
	"HOPE-backend/v2/models"
	"HOPE-backend/v2/services/auth/helper"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	sendblue "github.com/sendinblue/APIv3-go-library/lib"
)

type service struct {
	repo   models.AuthRepository
	mailer *sendblue.APIClient
}

func NewAuthService(repo models.AuthRepository, mailer *sendblue.APIClient) models.AuthService {
	return &service{
		repo:   repo,
		mailer: mailer,
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

func (ths *service) Register(req models.RegisterRequest) (*models.TokenPair, error) {
	hashedPassword, err := helper.EncryptPassword([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	newUser := models.DBUserWithProfile{
		Email:        req.Email,
		Password:     hashedPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		ProfilePhoto: req.ProfilePhoto,
		Weight:       req.Profile.Weight,
		Height:       req.Profile.Height,
		Job:          req.Profile.Job,
		Activities:   req.Profile.Activities,
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
	key, err := constructActivationKey(newUser)
	if err != nil {
		return nil, err
	}

	err = sendActivationKey(ths.mailer, key, newUser.Email)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func constructActivationKey(user models.DBUserWithProfile) (string, error) {
	userByte, err := json.Marshal(user)
	if err != nil {
		return "", errors.New("failed to construct activation key: " + err.Error())
	}

	key, err := helper.Encrypt(string(userByte))
	if err != nil {
		return "", errors.New("failed to encrypt activation key: " + err.Error())
	}

	return key, nil
}

func sendActivationKey(mailer *sendblue.APIClient, key, email string) error {
	// _, _, err := mailer.TransactionalEmailsApi.SendTransacEmail(
	// 	context.Background(),
	// 	sendblue.SendSmtpEmail{
	// 		Sender: &sendblue.SendSmtpEmailSender{
	// 			Name:  "dearhope",
	// 			Email: "no-reply@dearhope.id",
	// 		},
	// 		To: []sendblue.SendSmtpEmailTo{
	// 			{
	// 				Email: email,
	// 			},
	// 		},
	// 		Subject:     "Account Activation",
	// 		TextContent: "https://dearhope.id/activate?key=" + key,
	// 	},
	// )
	// if err != nil {
	// 	return errors.New("failed to send activation key: " + err.Error())
	// }

	fmt.Println("https://dearhope.id/activate?key=" + key)

	return nil
}

func (ths *service) GetLoggedInUser(id uint) (*models.UserResponse, error) {
	user, err := ths.repo.GetUserWithProfileByID(id)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		User: models.User{
			ID:           user.UserID,
			Email:        user.Email,
			Password:     "",
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			ProfilePhoto: user.ProfilePhoto,
		},
		Profile: models.Profile{
			ID:         user.ProfileID,
			Weight:     user.Weight,
			Height:     user.Height,
			Job:        user.Job,
			Activities: user.Activities,
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
		Weight:       req.Profile.Weight,
		Height:       req.Profile.Height,
		Job:          req.Profile.Job,
		Activities:   req.Profile.Activities,
	}

	updatedUser, err := ths.repo.UpdateUserWithProfile(newUser)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		User: models.User{
			ID:           updatedUser.UserID,
			Email:        updatedUser.Email,
			Password:     "",
			FirstName:    updatedUser.FirstName,
			LastName:     updatedUser.LastName,
			ProfilePhoto: updatedUser.ProfilePhoto,
		},
		Profile: models.Profile{
			ID:         updatedUser.ProfileID,
			Weight:     updatedUser.Weight,
			Height:     updatedUser.Height,
			Job:        updatedUser.Job,
			Activities: updatedUser.Activities,
		},
	}, nil
}

func (ths *service) Activate(req models.ActivateRequest) (*models.TokenPair, error) {
	userString, err := helper.Decrypt(strings.TrimPrefix(req.Key, "https://dearhope.id/activate?key="))
	if err != nil {
		return nil, errors.New("failed to decrypt activate key: " + err.Error())
	}

	var user models.DBUserWithProfile
	err = json.Unmarshal([]byte(userString), &user)
	if err != nil {
		return nil, errors.New("failed to activate account: " + err.Error())
	}

	user.IsActive = true
	_, err = ths.repo.UpdateUserWithProfile(user)
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(user.UserID, user.ProfileID, true)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}
