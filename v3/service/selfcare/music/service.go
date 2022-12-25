package music

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
)

type Service interface {
	List(f filter.ListMusic) ([]model.MusicPlaylistResponse, *model.ServiceError)
}
