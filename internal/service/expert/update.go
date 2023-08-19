package expert

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/expert"
	"HOPE-backend/internal/entity/response"
	"context"
	"database/sql"
	"errors"
	"net/http"
)

func (s *service) Update(ctx context.Context, req expert.CreateUpdateRequest) (bool, *response.ServiceError) {
	var (
		diff   []uint64
		svcErr *response.ServiceError
	)

	if len(req.TopicIds) > 0 {
		diff, svcErr = s.constructNewTopics(ctx, req.Id, req.TopicIds)
		if svcErr != nil {
			return false, svcErr
		}
	}

	_, err := s.repo.UpdateExpert(ctx, req.ToExpertModel(), req.TopicIds, diff)
	if err != nil {
		return false, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorUpdateExpertFailed,
			Err:  err,
		}
	}

	return true, nil
}

func (s *service) constructNewTopics(ctx context.Context, expertId uint64, topicIds []uint64) (
	[]uint64, *response.ServiceError) {
	var diff []uint64

	oldTopicIds, err := s.repo.GetExpertTopics(ctx, expertId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetExpertTopicIdsFailed,
			Err:  err,
		}
	}

	mb := make(map[uint64]struct{}, len(topicIds))
	for _, x := range topicIds {
		mb[x] = struct{}{}
	}
	for _, x := range oldTopicIds {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}

	return diff, nil
}
