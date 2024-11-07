package controllers

import (
	"net/http"
	"time"

	"github.com/dokuqui/monitor_scheduler/backend/models"
	"github.com/dokuqui/monitor_scheduler/backend/services"
	"github.com/gin-gonic/gin"
)

var script models.Script

func CreateScript(c *gin.Context) {
	if err := c.ShouldBindJSON(&script); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid script data"})
		return
	}

	owner := c.GetString("username")

	err := services.CreateScript(script.Name, script.Content, owner, script.UserGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create script"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Script created successfully"})
}

func ExecuteScript(c *gin.Context) {
	scriptID := c.Param("id")

	output, err := services.ExecuteScript(scriptID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute script"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"output": output})
}

func ScheduleScript(c *gin.Context) {
	scriptID := c.Param("id")
	var request struct {
		ScheduleTime time.Time `json:"schedule_time"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	err := services.ScheduleScript(scriptID, request.ScheduleTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to schedule script"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Script scheduled successfully"})
}

func UpdateScript(c *gin.Context) {
    var script models.Script
    if err := c.ShouldBindJSON(&script); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid script data"})
        return
    }

    err := services.UpdateScript(script)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update script"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Script updated successfully"})
}

func DeleteScript(c *gin.Context) {
	scriptID := c.Param("id")

	err := services.DeleteScript(scriptID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete script"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Script deleted successfully"})
}

func GetScriptByID(c *gin.Context) {
	scriptID := c.Param("id")

	script, err := services.GetScriptByID(scriptID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get script"})
		return
	}

	c.JSON(http.StatusOK, script)
}

func GetScriptsByUser(c *gin.Context) {
	owner := c.GetString("username")

	scripts, err := services.GetScriptsByUser(owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get scripts"})
		return
	}

	c.JSON(http.StatusOK, scripts)
}

func GetScriptsByUserGroup(c *gin.Context) {
	userGroup := c.GetString("user_group")

	scripts, err := services.GetScriptsByUserGroup(userGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get scripts"})
		return
	}

	c.JSON(http.StatusOK, scripts)
}

func GetAllScripts(c *gin.Context) {
	scripts, err := services.GetAllScripts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get scripts"})
		return
	}

	c.JSON(http.StatusOK, scripts)
}