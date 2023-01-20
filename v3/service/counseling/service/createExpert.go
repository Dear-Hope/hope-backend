package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
	"strings"
)

func (ths *service) CreateExpert(req model.CreateUpdateExpertRequest) (*model.ExpertResponse, *model.ServiceError) {
	if len(req.TopicIds) <= 0 {
		return nil, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ErrorExpertMinimalOneTopic),
		}
	}

	expert, err := ths.repo.StoreExpert(req.ToExpertModel(), req.TopicIds)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &model.ServiceError{
				Code: http.StatusBadRequest,
				Err:  errors.New(constant.ErrorExpertAlreadyExists),
			}
		} else {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  errors.New(constant.ErrorCreateExpertFailed),
			}
		}
	}

	return expert.ToExpertResponse(), nil
}
