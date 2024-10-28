package routes

import (
	"github.com/dokuqui/monitor_scheduler/backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.LoggingMiddleware())

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	config := cors.Config{
		AllowOrigins:     []string{allowedOrigins},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	r.Use(cors.New(config))

	UserRoutes(r)
	AdminRoutes(r)
	ManagerRoutes(r)

	return r
}
