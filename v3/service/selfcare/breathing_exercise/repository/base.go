package repository

import (
	"HOPE-backend/v3/service/selfcare/breathing_exercise"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) breathing_exercise.Repository {
	return &repository{
		db: db,
	}
}
