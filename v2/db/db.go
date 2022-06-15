package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgreSQLDatabase() *sqlx.DB {
	dsn := "host=localhost user=hope password=hope-database-pass dbname=hopev2 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	return sqlx.MustConnect("postgres", dsn)
}
