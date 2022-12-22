package repository

import (
	"HOPE-backend/v3/service/storyroom"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) storyroom.Repository {
	return &repository{
		db: db,
	}
}
