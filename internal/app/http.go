package app

import (
	"fmt"

	"HOPE-backend/config"
	"HOPE-backend/internal/server"
	_cache "HOPE-backend/pkg/cache"
	"HOPE-backend/pkg/db"
	_mailer "HOPE-backend/pkg/mailer"

	_authHandler "HOPE-backend/internal/api/auth"
	_expertHandler "HOPE-backend/internal/api/expert"
	_healthHandler "HOPE-backend/internal/api/health"
	_userHandler "HOPE-backend/internal/api/user"

	_expertRepo "HOPE-backend/internal/repository/expert"
	_scheduleRepo "HOPE-backend/internal/repository/schedule"
	_userRepo "HOPE-backend/internal/repository/user"

	_authService "HOPE-backend/internal/service/auth"
	_expertService "HOPE-backend/internal/service/expert"
	_scheduleService "HOPE-backend/internal/service/schedule"
	_userService "HOPE-backend/internal/service/user"
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
	expertRepo := _expertRepo.New(database)
	scheduleRepo := _scheduleRepo.New(database)

	// Init service
	authSvc := _authService.New(userRepo, expertRepo, mailer, cache)
	userSvc := _userService.New(userRepo, mailer, cache)
	expertSvc := _expertService.New(expertRepo, scheduleRepo)
	scheduleSvc := _scheduleService.New(scheduleRepo)

	// Init handler
	healthHandler := &_healthHandler.Handler{}
	authHandler := _authHandler.New(authSvc)
	userHandler := _userHandler.New(userSvc)
	expertHandler := _expertHandler.New(expertSvc, scheduleSvc)

	srv := server.Server{
		HealthHandler: healthHandler,
		AuthHandler:   authHandler,
		UserHandler:   userHandler,
		ExpertHandler: expertHandler,
	}

	return srv.Run(cfg.Server.Port, cfg.Server.ShutdownTimeoutInSec)
}
