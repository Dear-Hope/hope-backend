package db

import (
	"HOPE-backend/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgreSQLDatabase(config config.PostgreSQLConfig) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host,
		config.Username,
		config.Password,
		config.Name,
		config.Port,
		config.Sslmode,
		config.Timezone,
	)
	return sqlx.MustConnect("postgres", dsn)
}
