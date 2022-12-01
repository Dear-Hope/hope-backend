package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"net/http"
)

func (ths *service) GetByID(id uint) (*model.MovieResponse, *model.ServiceError) {
	movie, err := ths.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &model.ServiceError{
				Code: http.StatusNotFound,
				Err:  errors.New(constant.ERROR_MOVIE_NOT_FOUND),
			}
		}
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_MOVIE_FAILED),
		}
	}

	moods, err := ths.repo.GetAllMoodByMovieID(movie.ID)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_MOVIE_FAILED),
		}
	}

	needs, err := ths.repo.GetAllNeedByMovieID(movie.ID)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_MOVIE_FAILED),
		}
	}

	movie.Moods = moods
	movie.Needs = needs

	return movie.ToMovieResponse(), nil
}
