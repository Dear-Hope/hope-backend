package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (ths *service) Create(req model.MovieRequest) (*model.MovieResponse, *model.ServiceError) {
	if len(req.MoodIDs) <= 0 {
		return nil, &model.ServiceError{
			Code: http.StatusBadRequest,
			Err:  errors.New(constant.ERROR_MOVIE_MINIMAL_ONE_MOOD),
		}
	}

	newMovie := model.Movie{
		Title:       req.Title,
		Year:        req.Year,
		Country:     req.Country,
		Description: req.Description,
		TrailerLink: req.TrailerLink,
		PosterLink:  req.PosterLink,
		Genres:      req.Genres,
	}

	movie, err := ths.repo.Store(newMovie)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &model.ServiceError{
				Code: http.StatusBadRequest,
				Err:  errors.New(constant.ERROR_MOVIE_ALREADY_EXISTS),
			}
		} else {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  errors.New(constant.ERROR_CREATE_MOVIE_FAILED),
			}
		}
	}

	moods, err := ths.repo.StoreMoodDetail(movie.ID, req.MoodIDs)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf(constant.ERROR_STORE_MOVIE_DETAIL_FAILED, "mood"),
		}
	}

	needs, err := ths.repo.StoreNeedDetail(movie.ID, req.NeedIDs)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf(constant.ERROR_STORE_MOVIE_DETAIL_FAILED, "need"),
		}
	}

	movie.Moods = moods
	movie.Needs = needs

	return movie.ToMovieResponse(), nil
}
