package app

import (
	_authController "HOPE-backend/v3/service/auth/controller"
	_authRepo "HOPE-backend/v3/service/auth/repository"
	_authService "HOPE-backend/v3/service/auth/service"

	_moodController "HOPE-backend/v3/service/mood_tracker/controller"
	_moodRepo "HOPE-backend/v3/service/mood_tracker/repository"
	_moodService "HOPE-backend/v3/service/mood_tracker/service"

	_newsletterController "HOPE-backend/v3/service/newsletter/controller"
	_newsletterRepo "HOPE-backend/v3/service/newsletter/repository"
	_newsletterService "HOPE-backend/v3/service/newsletter/service"

	_movieController "HOPE-backend/v3/service/selfcare/movie/controller"
	_movieRepo "HOPE-backend/v3/service/selfcare/movie/repository"
	_movieService "HOPE-backend/v3/service/selfcare/movie/service"

	_categoryController "HOPE-backend/v3/service/selfcare/category/controller"
	_categoryRepo "HOPE-backend/v3/service/selfcare/category/repository"
	_categoryService "HOPE-backend/v3/service/selfcare/category/service"

	_breathingExerciseController "HOPE-backend/v3/service/selfcare/breathing_exercise/controller"
	_breathingExerciseRepo "HOPE-backend/v3/service/selfcare/breathing_exercise/repository"
	_breathingExerciseService "HOPE-backend/v3/service/selfcare/breathing_exercise/service"

	_selfHealingAudioController "HOPE-backend/v3/service/selfcare/self_healing_audio/controller"
	_selfHealingAudioRepo "HOPE-backend/v3/service/selfcare/self_healing_audio/repository"
	_selfHealingAudioService "HOPE-backend/v3/service/selfcare/self_healing_audio/service"

	_storyroomController "HOPE-backend/v3/service/storyroom/controller"
	_storyroomRepo "HOPE-backend/v3/service/storyroom/repository"
	_storyroomService "HOPE-backend/v3/service/storyroom/service"

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
	log.Println(config.MailerConfig.ApiKey)
	mailerCfg.AddDefaultHeader("partner-key", config.MailerConfig.PartnerKey)
	mailer := sendblue.NewAPIClient(mailerCfg)

	authRepo := _authRepo.NewRepository(database)
	authSvc := _authService.NewService(authRepo, mailer, cache)
	_authController.NewController(v3, authSvc)

	moodRepo := _moodRepo.NewRepository(database)
	moodSvc := _moodService.NewService(moodRepo, authRepo)
	_moodController.NewController(v3, moodSvc)

	newsletterRepo := _newsletterRepo.NewRepository(database)
	newsletterSvc := _newsletterService.NewService(newsletterRepo, mailer)
	_newsletterController.NewController(v3, newsletterSvc)

	movieRepo := _movieRepo.NewRepository(database)
	movieSvc := _movieService.NewService(movieRepo)
	_movieController.NewController(v3, movieSvc)

	categoryRepo := _categoryRepo.NewRepository(database)
	categorySvc := _categoryService.NewService(categoryRepo)
	_categoryController.NewController(v3, categorySvc)

	breathingExerciseRepo := _breathingExerciseRepo.NewRepository(database)
	breathingExerciseSvc := _breathingExerciseService.NewService(breathingExerciseRepo)
	_breathingExerciseController.NewController(v3, breathingExerciseSvc)

	selfHealingAudioRepo := _selfHealingAudioRepo.NewRepository(database)
	selfHealingAudioSvc := _selfHealingAudioService.NewService(selfHealingAudioRepo)
	_selfHealingAudioController.NewController(v3, selfHealingAudioSvc)

	storyroomRepo := _storyroomRepo.NewRepository(database)
	storyroomSvc := _storyroomService.NewService(storyroomRepo)
	_storyroomController.NewController(v3, storyroomSvc)

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
