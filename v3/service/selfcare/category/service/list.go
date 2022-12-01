package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
	"errors"
	"net/http"
)

func (ths *service) List(f filter.ListCategory) ([]model.CategoryResponse, *model.ServiceError) {
	var res []model.CategoryResponse
	categories, err := ths.repo.GetAll(f)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_LIST_CATEGORIES_FAILED),
		}
	}

	for _, category := range categories {
		total, err := ths.repo.GetCountCategory(constant.MAP_CATEGORY_TABLENAME[category.Name])
		if err != nil {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  errors.New(constant.ERROR_GET_TOTAL_PER_CATEGORY_FAILED),
			}
		}

		res = append(res, *category.ToCategoryResponse(total))
	}

	return res, nil
}
