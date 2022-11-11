package db

import (
	"HOPE-backend/config"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func RunDBMigrations(config config.PostgreSQLConfig, url string) error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Sslmode,
	)

	log.Println("Migrating base schema")
	migrateUp(dsn, url)

	files, err := ioutil.ReadDir("./v3/db/migrations")
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			log.Println("Migrating: " + f.Name())
			migrateUp(dsn+"&search_path="+f.Name(), url+f.Name())
		}
	}

	return nil
}

func migrateUp(dsn, url string) {
	migration, err := migrate.New(url, dsn)
	if err != nil {
		panic(err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}
