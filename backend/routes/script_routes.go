package routes

import (
	"github.com/dokuqui/monitor_scheduler/backend/controllers"
	"github.com/dokuqui/monitor_scheduler/backend/middleware"
	"github.com/gin-gonic/gin"
)

func ScriptRoutes(r *gin.Engine) {
    scripts := r.Group("/scripts")
    scripts.Use(middleware.AuthMiddleware())
    {
        scripts.POST("/create", controllers.CreateScript)
        scripts.POST("/execute/:id", controllers.ExecuteScript)
        scripts.POST("/schedule/:id", controllers.ScheduleScript)
        scripts.PUT("/update", controllers.UpdateScript)
        scripts.DELETE("/delete/:id", controllers.DeleteScript)
        scripts.GET("/get/:id", controllers.GetScriptByID)
        scripts.GET("/user", controllers.GetScriptsByUser)
        scripts.GET("/usergroup", middleware.RoleMiddleware("manager"), controllers.GetScriptsByUserGroup)
        scripts.GET("/all", middleware.RoleMiddleware("admin"), controllers.GetAllScripts)
        scripts.GET("/logs/:id", controllers.GetLogsByScriptID)
    }
}