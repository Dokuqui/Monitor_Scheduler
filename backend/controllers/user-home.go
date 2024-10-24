package controllers

import "github.com/gin-gonic/gin"

func UserHome(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome to your user dashboard!"})
}

func AdminDashboard(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome to your admin dashboard!"})
}

func ManagerDashboard(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome to your manager dashboard!"})
}
