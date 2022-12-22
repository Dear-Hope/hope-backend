package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths *service) ListCategories() ([]model.PostCategoryResponse, *model.ServiceError) {
	categories, err := ths.repo.GetAllCategories()
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_MOVIE_FAILED),
		}
	}

	return categories.ToListPostCategoryResponse(), nil
}
