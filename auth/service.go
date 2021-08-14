package auth

import (
	"HOPE-backend/auth/handler/helper"
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
		Email:    req.Email,
		Password: req.Password,
	}
	user, err := ths.repo.GetByEmailAndPassword(whereClause)
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(user.ID)
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

	newUser := &models.User{
		Email:    req.Email,
		Password: req.Password,
		IsAdmin:  false,
		Profile:  newUserProfile,
	}

	err := ths.repo.Create(newUser)
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(newUser.ID)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}
