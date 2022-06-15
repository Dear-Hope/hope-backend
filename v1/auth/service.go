package auth

import (
	"HOPE-backend/auth/helper"
	"HOPE-backend/v1/models"
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

	user, err := ths.repo.GetUserByEmail(whereClause)
	if err != nil {
		return nil, err
	}

	profile, err := ths.repo.GetProfileByUserID(user.Role, user.ID)
	if err != nil {
		return nil, err
	}

	err = helper.ComparePassword([]byte(req.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(user.ID, profile.GetProfileID())
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (ths *service) Register(req models.RegisterRequest) (*models.TokenPair, error) {
	var newUserProfile models.Profile
	if req.Role == "patient" {
		newUserProfile = &models.Patient{
			Weight:         req.PatientProfile.Weight,
			Height:         req.PatientProfile.Height,
			Job:            req.PatientProfile.Job,
			Activities:     req.PatientProfile.Activities,
			DiseaseHistory: req.PatientProfile.DiseaseHistory,
		}
	} else {
		newUserProfile = &models.Psychologist{
			Location:         req.PsychologistProfile.Location,
			Lisence:          req.PsychologistProfile.Lisence,
			MembershipNumber: req.PsychologistProfile.MembershipNumber,
			Experience:       req.PsychologistProfile.Experience,
			TopicCategory:    req.PsychologistProfile.TopicCategory,
			Specialization:   req.PsychologistProfile.Specialization,
		}
	}

	hashedPassword, err := helper.EncryptPassword([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
	}
	err = ths.repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	newUserProfile.SetUserID(newUser.ID)
	err = ths.repo.CreateProfile(newUser.Role, newUserProfile)
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(newUser.ID, newUserProfile.GetProfileID())
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (ths *service) GetLoggedInUser(id uint) (*models.UserResponse, error) {
	user, err := ths.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	profile, err := ths.repo.GetProfileByUserID(user.Role, user.ID)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{User: *user, Profile: profile}, nil
}

func (ths *service) UpdateLoggedInUser(req models.UpdateRequest) (*models.UserResponse, error) {
	newPassword, err := helper.EncryptPassword([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Email:     req.Email,
		Password:  newPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	newUser.ID = req.UserID
	if req.Password == "" {
		newUser.Password = req.Password
	}

	updatedUser, err := ths.repo.UpdateUser(newUser)
	if err != nil {
		return nil, err
	}

	var newUserProfile models.Profile
	if updatedUser.Role == "patient" {
		newUserProfile = &models.Patient{
			Weight:         req.PatientProfile.Weight,
			Height:         req.PatientProfile.Height,
			Job:            req.PatientProfile.Job,
			Activities:     req.PatientProfile.Activities,
			DiseaseHistory: req.PatientProfile.DiseaseHistory,
		}
	} else {
		newUserProfile = &models.Psychologist{
			Location:         req.PsychologistProfile.Location,
			Lisence:          req.PsychologistProfile.Lisence,
			MembershipNumber: req.PsychologistProfile.MembershipNumber,
			Experience:       req.PsychologistProfile.Experience,
			TopicCategory:    req.PsychologistProfile.TopicCategory,
			Specialization:   req.PsychologistProfile.Specialization,
		}
	}
	newUserProfile.SetUserID(updatedUser.ID)
	updatedUserProfile, err := ths.repo.UpdateProfile(updatedUser.Role, newUserProfile)
	if err != nil {
		return nil, err
	}

	updatedUser.Password = ""
	return &models.UserResponse{User: *updatedUser, Profile: updatedUserProfile}, nil
}

func (ths *service) ChangeOnlineStatus(userID uint) error {
	err := ths.repo.UpdateOnlineStatus(userID)
	if err != nil {
		return err
	}

	return nil
}
