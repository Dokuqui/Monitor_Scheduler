package routes

import (
	"github.com/dokuqui/monitor_scheduler/backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.LoggingMiddleware())

	UserRoutes(r)
	AdminRoutes(r)
	ManagerRoutes(r)

	return r
}
