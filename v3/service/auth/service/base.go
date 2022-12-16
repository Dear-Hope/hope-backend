package service

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/auth"
	"HOPE-backend/v3/service/auth/helper"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/pquerna/otp/totp"
	sendblue "github.com/sendinblue/APIv3-go-library/v2/lib"
)

type service struct {
	repo   auth.Repository
	mailer *sendblue.APIClient
	cache  *cache.Cache
}

func NewService(repo auth.Repository, mailer *sendblue.APIClient, cache *cache.Cache) auth.Service {
	return &service{
		repo:   repo,
		mailer: mailer,
		cache:  cache,
	}
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

func constructKey(user model.User) (string, error) {
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

func sendKey(mailer *sendblue.APIClient, template model.EmailTemplate) error {
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
			ScheduledAt: time.Now().Add(5 * time.Second),
		},
	)
	if err != nil {
		return errors.New("failed to send activation key: " + err.Error())
	}

	return nil
}
