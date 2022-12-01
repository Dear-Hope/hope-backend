package movie

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
)

type Repository interface {
	Store(newMovie model.Movie) (*model.Movie, error)
	StoreMoodDetail(id uint, moodIds []uint) (model.Moods, error)
	StoreNeedDetail(id uint, needIds []uint) (model.Needs, error)
	GetAll(f filter.ListMovie) (model.Movies, error)
	GetAllMoodByMovieID(movieID uint) (model.Moods, error)
	GetAllNeedByMovieID(movieID uint) (model.Needs, error)
	GetByID(id uint) (*model.Movie, error)
}
