package main

import (
	"developer-notes/config"
	"developer-notes/models"
	"log"
)

func main() {
	// Connect to database
	config.Connect()

	// Drop the notes table if it exists
	if err := config.DB.Migrator().DropTable("notes"); err != nil {
		log.Fatalf("Failed to drop notes table: %v", err)
	}

	// Create the notes table with the correct schema
	if err := config.DB.AutoMigrate(&models.Note{}); err != nil {
		log.Fatalf("Failed to create notes table: %v", err)
	}

	log.Println("Database cleaned and reset successfully!")
}
