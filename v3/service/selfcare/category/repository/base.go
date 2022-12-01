package repository

import (
	"HOPE-backend/v3/service/selfcare/category"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) category.Repository {
	return &repository{
		db: db,
	}
}
