package repository

import (
	"HOPE-backend/v3/service/newsletter"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) newsletter.Repository {
	return &repository{
		db: db,
	}
}
