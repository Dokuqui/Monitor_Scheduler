package services

import (
	"context"
	"log"
	"time"

	"github.com/dokuqui/monitor_scheduler/backend/db"
	"github.com/dokuqui/monitor_scheduler/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitializeUserCollection() {
	userCollection = db.Client.Database("monitorScheduler").Collection("user")
}

// CreateUser creates a new user in the db
func CreateUser(username, lastname, firstname, password, role string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, bson.M{
		"username":  username,
		"lastname":  lastname,
		"firstname": firstname,
		"password":  password,
		"role":      role,
	})

	if err != nil {
		log.Fatal(err)
	}
}

// GetUser retrieves a user by username from db
func GetUser(username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Return a nil User and a specific error if no user is found
			return models.User{}, nil // This indicates no user was found
		}
		return models.User{}, err // Return the actual error for other issues
	}

	return user, nil // Return the found user
}
