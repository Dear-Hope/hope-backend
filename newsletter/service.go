package newsletter

import (
	"HOPE-backend/models"
	"errors"
	"time"

	"github.com/beeker1121/mailchimp-go/lists/members"
)

type service struct {
	subsRepo models.NewsletterRepository
}

func NewNewsletterService(subsRepo models.NewsletterRepository) models.NewsletterService {
	return &service{
		subsRepo: subsRepo,
	}
}

func (ths *service) Subscribe(req models.NewSubscriberRequest) error {
	if req.Email == "" {
		return errors.New("invalid email given")
	}

	params := members.NewParams{
		EmailType:       "html",
		Status:          members.StatusSubscribed,
		Language:        "ID",
		TimestampSignup: time.UnixMilli(req.Time),
		EmailAddress:    req.Email,
	}

	_, err := members.New("ced232ba84", &params)
	if err != nil {
		return errors.New("error when subcribing to the newsletter")
	}

	newSubs := models.Subscription{
		Email:       req.Email,
		SubsribedAt: req.Time,
		// MemberID:    member.ID,
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

	err := ths.subsRepo.Delete(email)
	if err != nil {
		return err
	}

	return nil
}
