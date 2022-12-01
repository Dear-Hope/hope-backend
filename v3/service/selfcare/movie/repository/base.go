package repository

import (
	"HOPE-backend/v3/service/selfcare/movie"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) movie.Repository {
	return &repository{
		db: db,
	}
}
