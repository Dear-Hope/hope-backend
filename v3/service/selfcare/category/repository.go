package category

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
)

type Repository interface {
	GetAll(f filter.ListCategory) (model.Categories, error)
	GetCountCategory(tableName string) (int, error)
}
