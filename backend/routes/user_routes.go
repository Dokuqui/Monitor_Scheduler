package routes

import (
	"github.com/dokuqui/monitor_scheduler/backend/controllers"
	"github.com/dokuqui/monitor_scheduler/backend/middleware"
	"github.com/gin-gonic/gin"
)

// UserRoutes sets up the routes for user-related operations
func UserRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// Protected routes
	user := r.Group("/home")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/", controllers.UserHome) // This route requires authentication
	}
}

// AdminRoutes sets up the routes for admin-related operations
func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
	{
		admin.GET("/dashboard", controllers.AdminDashboard)
	}
}

// ManagerRoutes sets up the routes for manager-related operations
func ManagerRoutes(r *gin.Engine) {
	manager := r.Group("/manager")
	manager.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("manager"))
	{
		manager.GET("/dashboard", controllers.ManagerDashboard) // Fixed path with a leading slash
	}
}
