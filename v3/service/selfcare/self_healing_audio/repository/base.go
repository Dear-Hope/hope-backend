package repository

import (
	"HOPE-backend/v3/service/selfcare/self_healing_audio"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) self_healing_audio.Repository {
	return &repository{
		db: db,
	}
}
