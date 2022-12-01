package category

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
)

type Service interface {
	List(f filter.ListCategory) ([]model.CategoryResponse, *model.ServiceError)
}
