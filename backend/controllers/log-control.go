package controllers

import (
	"net/http"

	"github.com/dokuqui/monitor_scheduler/backend/services"
	"github.com/gin-gonic/gin"
)

func GetLogsByScriptID(c *gin.Context) {
	scriptID := c.Param("id")
	logs, err := services.GetLogsByScriptID(scriptID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get logs"})
		return
	}

	c.JSON(http.StatusOK, logs)
}