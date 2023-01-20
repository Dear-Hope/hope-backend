package repository

import (
	"HOPE-backend/v3/service/counseling"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) counseling.Repository {
	return &repository{
		db: db,
	}
}
