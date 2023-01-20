package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths *service) ListTopics() ([]model.TopicResponse, *model.ServiceError) {
	topics, err := ths.repo.GetAllTopics()
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ErrorGetListTopicsFailed),
		}
	}

	return topics.ToListTopicResponse(), nil
}
