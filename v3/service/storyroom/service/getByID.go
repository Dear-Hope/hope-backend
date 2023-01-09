package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"net/http"
)

func (ths *service) GetByID(id, userID uint) (*model.PostResponse, *model.ServiceError) {
	blockedUsers, err := ths.authRepo.GetBlockedUserByUserID(userID)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_BLOCKED_USER_FAILED),
		}
	}

	var excludedIDs []uint
	for _, user := range blockedUsers {
		excludedIDs = append(excludedIDs, user.BlockedUserID)
	}

	post, err := ths.repo.GetByID(id, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &model.ServiceError{
				Code: http.StatusNotFound,
				Err:  errors.New(constant.ERROR_POST_NOT_FOUND),
			}
		}
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_POST_FAILED),
		}
	}

	post.Comments, err = ths.repo.GetAllCommentByPostID(post.ID, excludedIDs)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_COMMENT_BY_POST_ID_FAILED),
		}
	}

	post.Categories, err = ths.repo.GetAllCategoryByPostID(post.ID)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_CATEGORY_BY_POST_ID_FAILED),
		}
	}

	post.TotalLikes = ths.repo.GetTotalLikesByPostID(post.ID)
	post.Liked = ths.repo.CheckIsPostLikedByUserID(post.ID, userID)

	return post.ToPostResponse(), nil
}
