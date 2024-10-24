package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dokuqui/monitor_scheduler/backend/models"
	"github.com/dokuqui/monitor_scheduler/backend/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secret")

var data struct {
	Username  string `json:"username"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

func Signup(c *gin.Context) {
	if c.BindJSON(&data) != nil || (data.Role != models.AdminRole && data.Role != models.ManagerRole && data.Role != models.UserRole) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	services.CreateUser(data.Username, data.Lastname, data.Firstname, string(hashedPassword), data.Role)

	c.JSON(http.StatusOK, gin.H{"message": "User created! Signup successful! Please login."})
}

func Login(c *gin.Context) {
	if c.BindJSON(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	user, err := services.GetUser(data.Username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

	switch user.Role {
	case "admin":
		c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin!", "redirect": "/admin/dashboard"})
	case "manager":
		c.JSON(http.StatusOK, gin.H{"message": "Welcome Manager!", "redirect": "/manager/dashboard"})
	case "user":
		c.JSON(http.StatusOK, gin.H{"message": "Welcome User!", "redirect": "/home"})
	}
}
