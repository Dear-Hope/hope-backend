package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"
)

func (ths *service) UpsertLike(postID, userID uint) *model.ServiceError {
	err := ths.repo.UpsertLike(postID, userID)
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf(constant.ERROR_UPSERT_LIKE_FAILED, postID),
		}
	}

	return nil
}
