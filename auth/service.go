package auth

import (
	"HOPE-backend/auth/helper"
	"HOPE-backend/models"
)

type service struct {
	repo models.AuthRepository
}

func NewAuthService(repo models.AuthRepository) models.AuthService {
	return &service{
		repo: repo,
	}
}

func (ths *service) Login(req models.LoginRequest) (*models.TokenPair, error) {
	whereClause := &models.User{
		Email: req.Email,
	}

	user, err := ths.repo.GetByEmail(whereClause)
	if err != nil {
		return nil, err
	}

	err = helper.ComparePassword([]byte(req.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(user.ID, user.Profile.ID)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (ths *service) Register(req models.RegisterRequest) (*models.TokenPair, error) {
	newUserProfile := models.UserProfile{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Weight:         req.Weight,
		Height:         req.Height,
		Job:            req.Job,
		Activities:     req.Activities,
		DiseaseHistory: req.DiseaseHistory,
	}

	hashedPassword, err := helper.EncryptPassword([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Email:    req.Email,
		Password: hashedPassword,
		IsAdmin:  false,
		Profile:  newUserProfile,
	}

	err = ths.repo.Create(newUser)
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(newUser.ID, newUser.Profile.ID)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (ths *service) GetLoggedInUser(id uint) (*models.User, error) {
	user, err := ths.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func (ths *service) UpdateLoggedInUser(req models.UpdateUserRequest) (*models.User, error) {
	newPassword, err := helper.EncryptPassword([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Email:    req.Email,
		Password: newPassword,
		IsAdmin:  req.IsAdmin,
		Profile: models.UserProfile{
			FirstName:      req.FirstName,
			LastName:       req.LastName,
			Weight:         req.Weight,
			Height:         req.Height,
			Job:            req.Job,
			Activities:     req.Activities,
			DiseaseHistory: req.DiseaseHistory,
			UserID:         req.UserID,
		},
	}

	newUser.ID = req.UserID
	newUser.Profile.ID = req.ProfileID
	if req.Password == "" {
		newUser.Password = req.Password
	}

	updatedUser, err := ths.repo.Update(newUser)
	if err != nil {
		return nil, err
	}

	updatedUser.Password = ""
	return updatedUser, nil
}
