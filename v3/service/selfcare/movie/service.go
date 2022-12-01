package movie

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
)

type Service interface {
	Create(req model.MovieRequest) (*model.MovieResponse, *model.ServiceError)
	List(f filter.ListMovie) ([]model.MovieResponse, *model.ServiceError)
	GetByID(id uint) (*model.MovieResponse, *model.ServiceError)
}
