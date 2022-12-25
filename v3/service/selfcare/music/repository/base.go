package repository

import (
	"HOPE-backend/v3/service/selfcare/music"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) music.Repository {
	return &repository{
		db: db,
	}
}
