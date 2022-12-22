package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths *service) DeletePostByID(postID, userID uint) *model.ServiceError {
	err := ths.repo.DeletePostByID(postID)
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_DELETE_POST_FAILED),
		}
	}

	err = ths.repo.DeleteCommentByPostID(postID)
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_DELETE_COMMENT_BY_POST_ID_FAILED),
		}
	}

	err = ths.repo.DeletePostCategoryByPostID(postID)
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_DELETE_CATEGORY_BY_POST_ID_FAILED),
		}
	}

	err = ths.repo.DeleteLikeByPostID(postID)
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_DELETE_LIKE_BY_POST_ID_FAILED),
		}
	}

	return nil
}
