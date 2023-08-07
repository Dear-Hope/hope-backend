package auth

import (
	"HOPE-backend/config"
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/user"
	"HOPE-backend/pkg/cache"
	"HOPE-backend/pkg/helpers"
	"HOPE-backend/pkg/mailer"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Repository interface {
	CreateUser(ctx context.Context, user user.User) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	VerifyUser(ctx context.Context, id uint64) error
	UpdateUser(ctx context.Context, user user.User) (*user.User, error)
}

type service struct {
	repo   Repository
	mailer mailer.Mailer
	cache  cache.Cache
}

func New(repo Repository, mailer mailer.Mailer, cache cache.Cache) *service {
	return &service{repo: repo, mailer: mailer, cache: cache}
}

func encryptPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("[AuthSvc][010001] error encrypting password: %v", err)
	}

	return string(hash), nil
}

func comparePassword(password, hash []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return fmt.Errorf("[AuthSvc][010002] error compare hash and password: %v", err)
	}

	return nil
}

func generateTokenPair(userId uint64, role string, verified bool) (*auth.TokenPairResponse, error) {
	access := jwt.New(jwt.SigningMethodHS256)

	atClaims := access.Claims.(jwt.MapClaims)
	atClaims["access"] = true
	atClaims["userId"] = userId
	atClaims["isVerified"] = verified
	atClaims["role"] = role
	atClaims["expires"] = time.Now().Add(config.Get().Jwt.AccessExpiryInHour * time.Hour).Unix()

	at, err := access.SignedString([]byte(config.Get().Server.SecretKey))
	if err != nil {
		return nil, fmt.Errorf("[AuthSvc][010003] error generate access token: %v", err)
	}

	refresh := jwt.New(jwt.SigningMethodHS256)

	rtClaims := refresh.Claims.(jwt.MapClaims)
	rtClaims["refresh"] = true
	rtClaims["userId"] = userId
	rtClaims["isVerified"] = verified
	rtClaims["role"] = role
	rtClaims["expires"] = time.Now().Add(config.Get().Jwt.RefreshExpiryInHour * time.Hour).Unix()

	rt, err := refresh.SignedString([]byte(config.Get().Server.SecretKey))
	if err != nil {
		return nil, fmt.Errorf("[AuthSvc][010004] error generate refresh token: %v", err)
	}

	return &auth.TokenPairResponse{
		Access:  at,
		Refresh: rt,
	}, nil
}

func generateSecretKey(email string) (string, error) {
	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      config.Get().Server.Name,
			AccountName: email,
		},
	)
	if err != nil {
		return "", fmt.Errorf("[AuthSvc][010005] error generate secret key: %v", err)
	}

	return key.Secret(), nil
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
