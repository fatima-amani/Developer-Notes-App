package config

import (
	"developer-notes/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Note{})
	DB = db
}