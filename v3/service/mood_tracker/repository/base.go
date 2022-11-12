package repository

import (
	"HOPE-backend/v3/service/mood_tracker"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) mood_tracker.Repository {
	return &repository{
		db: db,
	}
}
