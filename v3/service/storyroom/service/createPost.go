package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (ths *service) CreatePost(req model.PostRequest, authorID uint) (*model.PostResponse, *model.ServiceError) {
	if len(req.CategoryIDs) <= 0 {
		return nil, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ERROR_POST_MINIMAL_ONE_CATEGORY),
		}
	}

	newPost := model.Post{
		Content:  req.Content,
		AuthorID: authorID,
	}

	post, err := ths.repo.StorePost(newPost)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &model.ServiceError{
				Code: http.StatusBadRequest,
				Err:  errors.New(constant.ERROR_POST_ALREADY_EXISTS),
			}
		} else {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  errors.New(constant.ERROR_CREATE_POST_FAILED),
			}
		}
	}

	post.Categories, err = ths.repo.StorePostCategories(post.ID, req.CategoryIDs)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf(constant.ERROR_CREATE_CATEGORY_POST_FAILED, post.ID),
		}
	}

	return post.ToPostResponse(), nil
}
