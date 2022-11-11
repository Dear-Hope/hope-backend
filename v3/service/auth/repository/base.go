package repository

import (
	"HOPE-backend/v3/service/auth"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) auth.Repository {
	return &repository{
		db: db,
	}
}
