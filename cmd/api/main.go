package main

import (
	"log"

	"github.com/Bantamlak12/personal_book_library_manager/internal/api"
	"github.com/Bantamlak12/personal_book_library_manager/internal/config"
	"github.com/Bantamlak12/personal_book_library_manager/internal/repository"
)

func main() {
	// Initialize the database
	repo := repository.NewSQLiteRepository()
	log.Println("Database initialized successfully.")
	defer repo.CloseDB()

	// Load the configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v\n", err)
	}

	// Initialize Gin router
	server := api.SetupRouter(repo)

	// Start the server
	log.Printf("Server is running on %s\n", cfg.ServerAddress)
	if err := server.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
