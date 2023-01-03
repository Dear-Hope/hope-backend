package auth

import "HOPE-backend/v3/model"

type Service interface {
	Login(model.LoginRequest) (*model.TokenPairResponse, *model.ServiceError)
	Register(model.RegisterRequest) (*model.TokenPairResponse, *model.ServiceError)
	GetLoggedInUser(uint) (*model.UserResponse, *model.ServiceError)
	UpdateLoggedInUser(model.UpdateRequest) (*model.UserResponse, *model.ServiceError)
	Activate(model.ActivateRequest) (*model.TokenPairResponse, *model.ServiceError)
	ResetPassword(model.ResetPasswordRequest) *model.ServiceError
	ChangePassword(model.ChangePasswordRequest) (*model.TokenPairResponse, *model.ServiceError)
	SaveProfilePhoto(model.SaveProfilePhotoRequest) (string, *model.ServiceError)
	ResendActivationCode(model.ResetPasswordRequest) *model.ServiceError
	DeleteUser(model.ResetPasswordRequest) *model.ServiceError
	BlockUser(userID uint, req model.BlockUserRequest) *model.ServiceError
}
