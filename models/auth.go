package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"unique;not null"`
	Password     string `json:"password,omitempty" gorm:"not null"`
	Role         string `json:"role" gorm:"not null;not blank"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	ProfilePhoto string `json:"profile_photo"`
}

type Profile interface {
	GetProfileID() uint
	SetUserID(id uint)
}

type Patient struct {
	gorm.Model
	Weight         float32 `json:"weight" gorm:"not null"`
	Height         float32 `json:"height" gorm:"not null"`
	Job            string  `json:"job"`
	Activities     string  `json:"activities"`
	DiseaseHistory string  `json:"disease_history"`
	UserID         uint    `json:"user_id"`
}

type Psychologist struct {
	gorm.Model
	Location         string `json:"location"`
	Lisence          string `json:"lisence"`
	MembershipNumber int    `json:"membership_number"`
	Experience       string `json:"experience"`
	TopicCategory    string `json:"topic_category"`
	Specialization   string `json:"specialization"`
	UserID           uint   `json:"user_id"`
}

func (User) TableName() string {
	return "user"
}

func (Patient) TableName() string {
	return "patient"
}

func (Psychologist) TableName() string {
	return "psychologist"
}

func (ths *Patient) GetProfileID() uint {
	return ths.ID
}
func (ths *Patient) SetUserID(id uint) {
	ths.UserID = id
}
func (ths *Psychologist) GetProfileID() uint {
	return ths.ID
}
func (ths *Psychologist) SetUserID(id uint) {
	ths.UserID = id
}

type AuthService interface {
	Login(LoginRequest) (*TokenPair, error)
	Register(RegisterRequest) (*TokenPair, error)
	GetLoggedInUser(uint) (*UserResponse, error)
	UpdateLoggedInUser(UpdateRequest) (*UserResponse, error)
}

type AuthRepository interface {
	CreateUser(*User) error
	CreateProfile(string, Profile) error
	GetUserByEmail(*User) (*User, error)
	GetUserByID(uint) (*User, error)
	GetProfileByUserID(string, uint) (Profile, error)
	UpdateUser(*User) (*User, error)
	UpdateProfile(string, Profile) (Profile, error)
}

type UserResponse struct {
	User
	Profile Profile `json:"profile,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email               string                      `json:"email"`
	Password            string                      `json:"password"`
	FirstName           string                      `json:"first_name,omitempty"`
	LastName            string                      `json:"last_name,omitempty"`
	Role                string                      `json:"role"`
	ProfilePhoto        string                      `json:"profile_photo,omitempty"`
	PatientProfile      RegisterPatientRequest      `json:"patient_profile,omitempty"`
	PsychologistProfile RegisterPsychologistRequest `json:"psychologist_profile,omitempty"`
}

type RegisterPatientRequest struct {
	Weight         float32 `json:"weight"`
	Height         float32 `json:"height"`
	Job            string  `json:"job,omitempty"`
	Activities     string  `json:"activities,omitempty"`
	DiseaseHistory string  `json:"disease_history,omitempty"`
}

type RegisterPsychologistRequest struct {
	Location         string `json:"location,omitempty"`
	Lisence          string `json:"lisence,omitempty"`
	MembershipNumber int    `json:"membership_number,omitempty"`
	Experience       string `json:"experience,omitempty"`
	TopicCategory    string `json:"topic_category,omitempty"`
	Specialization   string `json:"specialization,omitempty"`
}

type UpdateRequest struct {
	Email               string                    `json:"email,omitempty"`
	Password            string                    `json:"password,omitempty"`
	IsAdmin             bool                      `json:"is_admin,omitempty"`
	FirstName           string                    `json:"first_name,omitempty"`
	LastName            string                    `json:"last_name,omitempty"`
	PhotoProfile        string                    `json:"photo_profile,omitempty"`
	UserID              uint                      `json:"user_id,omitempty"`
	ProfileID           uint                      `json:"profile_id,omitempty"`
	PatientProfile      UpdatePatientRequest      `json:"patient_profile,omitempty"`
	PsychologistProfile UpdatePsychologistRequest `json:"psychologist_profile,omitempty"`
}

type UpdatePatientRequest struct {
	Weight         float32 `json:"weight,omitempty"`
	Height         float32 `json:"height,omitempty"`
	Job            string  `json:"job,omitempty"`
	Activities     string  `json:"activities,omitempty"`
	DiseaseHistory string  `json:"disease_history,omitempty"`
}

type UpdatePsychologistRequest struct {
	Location         string `json:"location,omitempty"`
	Lisence          string `json:"lisence,omitempty"`
	MembershipNumber int    `json:"membership_number,omitempty"`
	Experience       string `json:"experience,omitempty"`
	TopicCategory    string `json:"topic_category,omitempty"`
	Specialization   string `json:"specialization,omitempty"`
}

type TokenPair struct {
	Access  string `json:"access,omitempty"`
	Refresh string `json:"refresh,omitempty"`
}
