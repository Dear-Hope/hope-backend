package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
	"errors"
	"net/http"
)

func (ths *service) List(f filter.ListMusic) ([]model.MusicPlaylistResponse, *model.ServiceError) {
	playlists, err := ths.repo.GetAll(f)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_LIST_MUSIC_PLAYLIST_FAILED),
		}
	}

	return playlists.ToListMusicPlaylistResponse(), nil
}
