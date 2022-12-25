package music

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
)

type Repository interface {
	GetAll(f filter.ListMusic) (model.MusicPlaylists, error)
}
