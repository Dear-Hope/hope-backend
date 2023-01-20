package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"net/http"
)

func (ths *service) UpdateExpert(req model.CreateUpdateExpertRequest) (*model.ExpertResponse, *model.ServiceError) {
	oldTopicIds, err := ths.repo.GetExpertTopics(req.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ErrorGetExpertTopicIdsFailed),
		}
	}

	mb := make(map[uint]struct{}, len(req.TopicIds))
	for _, x := range req.TopicIds {
		mb[x] = struct{}{}
	}
	var diff []uint
	for _, x := range oldTopicIds {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}

	updatedExpert, err := ths.repo.UpdateExpert(req.ToExpertModel(), req.TopicIds, diff)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ErrorGetExpertTopicIdsFailed),
		}
	}

	return updatedExpert.ToExpertResponse(), nil
}
