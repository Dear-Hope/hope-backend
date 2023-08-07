package app

import (
	"HOPE-backend/config"
	_authHandler "HOPE-backend/internal/api/auth"
	"HOPE-backend/internal/api/health"
	_userHandler "HOPE-backend/internal/api/user"
	_userRepo "HOPE-backend/internal/repository/user"
	"HOPE-backend/internal/server"
	_authService "HOPE-backend/internal/service/auth"
	_userService "HOPE-backend/internal/service/user"
	_cache "HOPE-backend/pkg/cache"
	"HOPE-backend/pkg/db"
	_mailer "HOPE-backend/pkg/mailer"
	"fmt"
)

func Init(cfg *config.Config) error {
	// Init db
	database := db.NewPostgresDatabase(cfg.Database.Postgres)
	defer func() {
		_ = database.Close()
	}()

	// Run migrations
	// TODO: move run migrations to dedicated service for migration
	if cfg.FeatureFlag.RunMigrations {
		if err := db.RunDBMigrations(cfg.Database.Postgres, cfg.MigrationFileUrl); err != nil {
			return fmt.Errorf("failed to migrate database: %v", err)
		}
	}

	// Init cache
	cache := _cache.New(cfg.Cache)

	// Init mailer
	mailer := _mailer.New(cfg.Mailer)

	// Init repository
	userRepo := _userRepo.New(database)

	// Init service
	authSvc := _authService.New(userRepo, mailer, cache)
	userSvc := _userService.New(userRepo)

	// Init handler
	healthHandler := &health.Handler{}
	authHandler := _authHandler.New(authSvc)
	userHandler := _userHandler.New(userSvc)

	srv := server.Server{
		HealthHandler: healthHandler,
		AuthHandler:   authHandler,
		UserHandler:   userHandler,
	}

	return srv.Run(cfg.Server.Port, cfg.Server.ShutdownTimeoutInSec)
}
