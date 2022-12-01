package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	sendblue "github.com/sendinblue/APIv3-go-library/v2/lib"
)

func (ths *service) Subscribe(req model.NewSubscriberRequest) *model.ServiceError {
	if req.Email == "" {
		return &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ERROR_INVALID_EMAIL),
		}
	}

	log.Println(req)
	_, _, err := ths.mailer.ContactsApi.CreateContact(
		context.Background(),
		sendblue.CreateContact{
			Email:   req.Email,
			ListIds: []int64{3},
		},
	)
	if err != nil && !strings.Contains(err.Error(), "400 Bad Request") {
		log.Printf("failed to subscribe: %s", err.Error())
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_SUBSCRIBE_FAILED),
		}
	}

	newSubs := model.Subscription{
		Email:        req.Email,
		SubscribedAt: req.Time,
		// MemberID:    "placeholder for now",
	}

	err = ths.subsRepo.Create(newSubs)
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_SUBSCRIBE_FAILED),
		}
	}

	return nil
}