package app

import (
	"HOPE-backend/config"
	"HOPE-backend/infra/db"
	"HOPE-backend/internal/api/health"
	"HOPE-backend/internal/server"
	"fmt"
)

func Init(cfg *config.Config) error {
	database := db.NewPostgresDatabase(cfg.Database)
	defer func() {
		_ = database.Close()
	}()

	// Run migrations
	// TODO: move run migrations to dedicated service for migration
	if cfg.FeatureFlag.RunMigrations {
		if err := db.RunDBMigrations(cfg.Database, cfg.MigrationFileUrl); err != nil {
			return fmt.Errorf("failed to migrate database: %v", err)
		}
	}

	// Init handler
	healthHandler := &health.Handler{}

	srv := server.Server{
		HealthHandler: healthHandler,
	}

	return srv.Run(cfg.Server.Port, cfg.Server.ShutdownTimeoutInSec)
}
