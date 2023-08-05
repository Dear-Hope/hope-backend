package app

import (
	"HOPE-backend/config"
	_authHandler "HOPE-backend/internal/api/auth"
	"HOPE-backend/internal/api/health"
	_authRepo "HOPE-backend/internal/repository/auth"
	"HOPE-backend/internal/server"
	_authService "HOPE-backend/internal/service/auth"
	"HOPE-backend/pkg/db"
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

	// Init repository
	authRepo := _authRepo.New(database)

	// Init service
	authSvc := _authService.New(authRepo)

	// Init handler
	healthHandler := &health.Handler{}
	authHandler := _authHandler.New(authSvc)

	srv := server.Server{
		HealthHandler: healthHandler,
		AuthHandler:   authHandler,
	}

	return srv.Run(cfg.Server.Port, cfg.Server.ShutdownTimeoutInSec)
}
