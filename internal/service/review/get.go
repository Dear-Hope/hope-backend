package review

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/review"
	"context"
	"net/http"
)

func (s *service) Get(ctx context.Context, expertId uint64) (*review.ListResponse, *response.ServiceError) {
	reviews, err := s.repo.GetReviewsByExpertId(ctx, expertId)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetExpertReviewFailed,
			Err:  err,
		}
	}

	return reviews.ToListResponse(), nil
}
