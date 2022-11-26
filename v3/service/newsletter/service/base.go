package service

import (
	"HOPE-backend/v3/service/newsletter"

	sendblue "github.com/sendinblue/APIv3-go-library/v2/lib"
)

type service struct {
	subsRepo newsletter.Repository
	mailer   *sendblue.APIClient
}

func NewService(subsRepo newsletter.Repository, mailer *sendblue.APIClient) newsletter.Service {
	return &service{
		subsRepo: subsRepo,
		mailer:   mailer,
	}
}
