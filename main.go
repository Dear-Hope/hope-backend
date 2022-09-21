package main

import (
	"HOPE-backend/v2/services/auth"
	_authHandler "HOPE-backend/v2/services/auth/handler"
	_authRepo "HOPE-backend/v2/services/auth/repository"
	"HOPE-backend/v2/services/moodtracker"
	_moodHandler "HOPE-backend/v2/services/moodtracker/handler"
	_moodRepo "HOPE-backend/v2/services/moodtracker/repository"
	"HOPE-backend/v2/services/newsletter"
	_newsletterHandler "HOPE-backend/v2/services/newsletter/handler"
	_newsletterRepo "HOPE-backend/v2/services/newsletter/repository"
	"HOPE-backend/v2/services/selfcare"
	_selfCareHandler "HOPE-backend/v2/services/selfcare/handler"
	_selfCareRepo "HOPE-backend/v2/services/selfcare/repository"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"HOPE-backend/config"
	"HOPE-backend/v2/db"

	sendblue "github.com/sendinblue/APIv3-go-library/lib"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalf("Failed to load app configuration: %s", err)
	}
	router := newRouter()

	database := db.NewPostgreSQLDatabase(config.DBConfig)
	if err := db.RunDBMigrations(config.DBConfig, config.MigrationFileURL); err != nil {
		log.Fatalf("failed to migrate database: %s", err)
	}

	router.Static("/assets", "./assets")

	router.GET("/server/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is OK!")
	})

	v2 := router.Group("/api/v2")

	mailerCfg := sendblue.NewConfiguration()
	mailerCfg.AddDefaultHeader("api-key", config.MailerConfig.ApiKey)
	mailerCfg.AddDefaultHeader("partner-key", config.MailerConfig.PartnerKey)
	mailer := sendblue.NewAPIClient(mailerCfg)

	authRepo := _authRepo.NewPostgreSQLRepository(database)
	authSvc := auth.NewAuthService(authRepo, mailer)
	_authHandler.NewAuthHandler(v2, authSvc)

	moodRepo := _moodRepo.NewPostgreSQLRepository(database)
	moodSvc := moodtracker.NewMoodTrackerService(moodRepo, authRepo)
	_moodHandler.NewMoodTrackerHandler(v2, moodSvc)

	selfCareRepo := _selfCareRepo.NewPostgreSQLRepository(database)
	selfCareSvc := selfcare.NewSelfCareService(selfCareRepo)
	_selfCareHandler.NewSelfCareHandler(v2, selfCareSvc)

	newsletterRepo := _newsletterRepo.NewPostgreSQLRepository(database)
	newsletterSvc := newsletter.NewNewsletterService(newsletterRepo, mailer)
	_newsletterHandler.NewNewsletterService(v2, newsletterSvc)

	router.Run(":8000")
}

func newRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("*")
	config.AddAllowMethods("OPTIONS")
	router.Use(cors.New(config))

	return router
}
