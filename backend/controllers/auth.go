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
	Manager   string `json:"manager,omitempty"`
}

func determineRedirect(role string) string {
	switch role {
	case "admin":
		return "/admin/dashboard"
	case "manager":
		return "/manager/dashboard"
	case "user":
		return "/home"
	default:
		return "/login" // Fallback in case of an undefined role
	}
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

	services.AuthUser(data.Username, data.Lastname, data.Firstname, string(hashedPassword), data.Role)

	c.JSON(http.StatusOK, gin.H{"message": "User created! Signup successful! Please login."})
}

func Login(c *gin.Context) {
	// Bind JSON data
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Retrieve the user from the database
	user, err := services.GetUserByUsername(data.Username)
	if err != nil {
		log.Println("Error retrieving user:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the provided password with the stored hashed password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create a new JWT token with HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Println("Error signing token:", err) // Log the error for debugging
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Prepare the response based on the user's role
	response := gin.H{
		"token":    tokenString,
		"message":  "Welcome " + user.Role + "!",
		"role":     user.Role,                    // Ensure this line exists
		"redirect": determineRedirect(user.Role), // Call to the determineRedirect function
	}

	// Send the response
	c.JSON(http.StatusOK, response)
}
