package auth

import (
	"HOPE-backend/internal/entity/expert"
	"HOPE-backend/internal/entity/user"
	"HOPE-backend/pkg/cache"
	"HOPE-backend/pkg/helpers"
	"HOPE-backend/pkg/mailer"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	UpdateUser(ctx context.Context, user user.User) (*user.User, error)
}

type ExpertRepository interface {
	GetExpertByEmail(ctx context.Context, email string) (*expert.Expert, error)
	//UpdateUser(ctx context.Context, user user.User) (*user.User, error)
}

type service struct {
	userRepo   UserRepository
	expertRepo ExpertRepository
	mailer     mailer.Mailer
	cache      cache.Cache
}

func New(uRepo UserRepository, eRepo ExpertRepository, mailer mailer.Mailer, cache cache.Cache) *service {
	return &service{userRepo: uRepo, expertRepo: eRepo, mailer: mailer, cache: cache}
}

func comparePassword(password, hash []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return fmt.Errorf("[AuthSvc][010002] error compare hash and password: %v", err)
	}

	return nil
}

func constructKey(user user.User) (string, error) {
	userByte, err := json.Marshal(user)
	if err != nil {
		return "", fmt.Errorf("[AuthSvc][010008] failed to construct key: %v", err)
	}

	key, err := helpers.Encrypt(string(userByte))
	if err != nil {
		return "", fmt.Errorf("[AuthSvc][010009] failed to encrypt key: %v", err)
	}

	return key, nil
}
