package newsletter

import (
	"HOPE-backend/v2/models"
	"context"
	"errors"

	sendblue "github.com/sendinblue/APIv3-go-library/lib"
)

type service struct {
	subsRepo models.NewsletterRepository
	mailer   *sendblue.APIClient
}

func NewNewsletterService(subsRepo models.NewsletterRepository, mailer *sendblue.APIClient) models.NewsletterService {
	return &service{
		subsRepo: subsRepo,
		mailer:   mailer,
	}
}

func (ths *service) Subscribe(req models.NewSubscriberRequest) error {
	if req.Email == "" {
		return errors.New("invalid email given")
	}

	_, _, err := ths.mailer.ContactsApi.CreateContact(
		context.Background(),
		sendblue.CreateContact{
			Email:   req.Email,
			ListIds: []int64{3},
		},
	)
	if err != nil {
		return errors.New("failed to subscribe: " + err.Error())
	}

	newSubs := models.Subscription{
		Email:        req.Email,
		SubscribedAt: req.Time,
		// MemberID:    "placeholder for now",
	}

	err = ths.subsRepo.Create(newSubs)
	if err != nil {
		return err
	}

	return nil
}

func (ths *service) Unsubscribe(email string) error {
	if email == "" {
		return errors.New("invalid email given")
	}

	_, err := ths.mailer.ContactsApi.DeleteContact(
		context.Background(),
		email,
	)
	if err != nil {
		return errors.New("failed to unsubscribe: " + err.Error())
	}

	err = ths.subsRepo.Delete(email)
	if err != nil {
		return err
	}

	return nil
}
