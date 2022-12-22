package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/storyroom/filter"
	"errors"
	"net/http"
)

func (ths *service) List(f filter.List, userID uint) ([]model.PostResponse, *model.ServiceError) {
	posts, err := ths.repo.GetAll(f)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_LIST_POST_FAILED),
		}
	}

	for i, post := range posts {
		posts[i].Categories, err = ths.repo.GetAllCategoryByPostID(post.ID)
		if err != nil {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  errors.New(constant.ERROR_GET_CATEGORY_BY_POST_ID_FAILED),
			}
		}

		posts[i].TotalLikes = ths.repo.GetTotalLikesByPostID(post.ID)
		posts[i].TotalComments = ths.repo.GetTotalCommentsByPostID(post.ID)
		posts[i].Liked = ths.repo.CheckIsPostLikedByUserID(post.ID, userID)
	}

	return posts.ToListPostResponse(), nil
}
