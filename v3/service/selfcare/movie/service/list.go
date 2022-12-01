package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
	"errors"
	"net/http"
)

func (ths *service) List(f filter.ListMovie) ([]model.MovieResponse, *model.ServiceError) {
	movies, err := ths.repo.GetAll(f)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_LIST_MOVIE_FAILED),
		}
	}

	for i, movie := range movies {
		moods, err := ths.repo.GetAllMoodByMovieID(movie.ID)
		if err != nil {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  errors.New(constant.ERROR_GET_LIST_MOVIE_FAILED),
			}
		}

		needs, err := ths.repo.GetAllNeedByMovieID(movie.ID)
		if err != nil {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  errors.New(constant.ERROR_GET_LIST_MOVIE_FAILED),
			}
		}

		movies[i].Moods = moods
		movies[i].Needs = needs
	}

	return movies.ToListMovieResponse(), nil
}
