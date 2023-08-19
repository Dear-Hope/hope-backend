package main

import (
	"HOPE-backend/config"
	"HOPE-backend/internal/app"
	"log"
)

func main() {
	var cfg *config.Config

	// load config
	err := config.Load(
		config.WithConfigFile("config"),
		config.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	cfg = config.Get()
	log.Println(cfg)
	log.Println(cfg.Database)

	log.Println("starting dear hope backend service... ")
	if err := app.Init(cfg); err != nil {
		log.Fatal(err)
	}
}
