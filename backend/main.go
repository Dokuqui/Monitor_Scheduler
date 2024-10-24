package main

import (
	"log"

	"github.com/dokuqui/monitor_scheduler/backend/db"
	"github.com/dokuqui/monitor_scheduler/backend/routes"
)

func main() {
	db.Connect()

	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
