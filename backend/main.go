package main

import (
	"github.com/dokuqui/monitor_scheduler/backend/db"
	"github.com/dokuqui/monitor_scheduler/backend/routes"
	"github.com/dokuqui/monitor_scheduler/backend/services"
	"log"
)

func main() {
	db.Connect()
	// Initialize the user collection
	services.InitializeUserCollection()

	// Set up routes
	r := routes.SetupRouter()

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
