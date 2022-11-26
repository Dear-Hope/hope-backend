package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"context"
	"errors"
	"log"
	"net/http"
)

func (ths *service) Unsubscribe(email string) *model.ServiceError {
	if email == "" {
		return &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ERROR_INVALID_EMAIL),
		}
	}

	_, err := ths.mailer.ContactsApi.DeleteContact(
		context.Background(),
		email,
	)
	if err != nil {
		log.Printf("failed to unsubscribe: %s", err.Error())
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_UNSUBSCRIBE_FAILED),
		}
	}

	err = ths.subsRepo.Delete(email)
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_UNSUBSCRIBE_FAILED),
		}
	}

	return nil
}
