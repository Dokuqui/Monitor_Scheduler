package controllers

import (
	"github.com/dokuqui/monitor_scheduler/backend/models"
	"github.com/dokuqui/monitor_scheduler/backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to create a user"})
		return
	}

	err := services.CreateUser(user.Username, user.Lastname, user.Firstname, user.Password, user.Role, user.Manager)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	// Check role
	role := c.GetString("role")
	if role != "admin" && role != "manager" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this user"})
		return
	}

	// Update user
	err := services.UpdateUser(user.Username, user.Lastname, user.Firstname, user.Password, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	username := c.Param("username")

	// Ensure the user is an Admin
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this user"})
		return
	}

	err := services.DeleteUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func ListUsers(c *gin.Context) {
	role := c.GetString("role")

	var users []models.User
	var err error

	if role == "admin" {
		users, err = services.ListAllUsers()
	} else if role == "manager" {
		username := c.GetString("username")
		users, err = services.ListUsersByManager(username)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func UpdateOwnCredentials(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	username := c.GetString("username")
	err := services.UpdateUser(username, user.Lastname, user.Firstname, user.Password, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Credentials updated successfully"})
}