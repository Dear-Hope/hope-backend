package app

import (
	_authController "HOPE-backend/v3/service/auth/controller"
	_authRepo "HOPE-backend/v3/service/auth/repository"
	_authService "HOPE-backend/v3/service/auth/service"
	"log"
	"net/http"

	"HOPE-backend/config"
	"HOPE-backend/v3/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	sendblue "github.com/sendinblue/APIv3-go-library/lib"
)

func Start() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalf("Failed to load app configuration: %s", err)
	}
	router := newRouter()

	database := db.NewPostgreSQLDatabase(config.DBConfig)
	if err := db.RunDBMigrations(config.DBConfig, config.MigrationFileURL); err != nil {
		log.Fatalf("failed to migrate database: %s", err)
	}

	cache := db.NewInmemCache(config.CacheConfig)

	router.Static("/assets", "./assets")

	router.GET("/server/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Server is OK!")
	})

	v3 := router.Group("/api/v3")

	mailerCfg := sendblue.NewConfiguration()
	mailerCfg.AddDefaultHeader("api-key", config.MailerConfig.ApiKey)
	mailerCfg.AddDefaultHeader("partner-key", config.MailerConfig.PartnerKey)
	mailer := sendblue.NewAPIClient(mailerCfg)

	authRepo := _authRepo.NewRepository(database)
	authSvc := _authService.NewService(authRepo, mailer, cache)
	_authController.NewController(v3, authSvc)

	router.Logger.Fatal(router.Start(":8000"))
}

func newRouter() *echo.Echo {
	router := echo.New()
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: middleware.DefaultCORSConfig.AllowOrigins,
		AllowHeaders: []string{"*"},
		AllowMethods: append(middleware.DefaultCORSConfig.AllowMethods, http.MethodOptions),
	}))

	router.Use(middleware.Logger())

	return router
}
