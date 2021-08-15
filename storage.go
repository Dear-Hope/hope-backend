package main

import (
	"HOPE-backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQLDatabase() *gorm.DB {
	dsn := "host=localhost user=hope password=hope-database-pass dbname=hope port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("could not open database connection")
	}

	migrateTable(db)

	return db
}

func migrateTable(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserProfile{})
}