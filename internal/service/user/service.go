package user

import (
	"HOPE-backend/config"
	"HOPE-backend/internal/entity/user"
	"HOPE-backend/pkg/cache"
	"HOPE-backend/pkg/mailer"
	"context"
	"fmt"
	"github.com/pquerna/otp/totp"
)

type repository interface {
	CreateUser(ctx context.Context, user user.User) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	GetUserById(ctx context.Context, uid uint64) (*user.User, error)
	UpdateUser(ctx context.Context, user user.User) (*user.User, error)
	VerifyUser(ctx context.Context, id uint64) error
}

type service struct {
	repo   repository
	mailer mailer.Mailer
	cache  cache.Cache
}

func New(repo repository, mailer mailer.Mailer, cache cache.Cache) *service {
	return &service{repo: repo, mailer: mailer, cache: cache}
}

func generateSecretKey(email string) (string, error) {
	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      config.Get().Server.Name,
			AccountName: email,
		},
	)
	if err != nil {
		return "", fmt.Errorf("[UserSvc] error generate secret key: %v", err)
	}

	return key.Secret(), nil
}
