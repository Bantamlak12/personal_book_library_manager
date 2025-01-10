package main

import (
	"log"

	"github.com/Bantamlak12/personal_book_library_manager/internal/api"
	"github.com/Bantamlak12/personal_book_library_manager/internal/config"
)

func main() {
	// Load the configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v\n", err)
	}

	// Initialize Gin router
	server := api.SetupRouter()

	// Start the server
	log.Printf("Server is running on %s\n", cfg.ServerAddress)
	if err := server.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
