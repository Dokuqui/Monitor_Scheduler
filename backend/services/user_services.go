package services

import (
	"context"
	"golang.org/x/crypto/bcrypt"
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

func AuthUser(username, lastname, firstname, password, role string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()

	_, err := userCollection.InsertOne(ctx, bson.M{
		"username":   username,
		"lastname":   lastname,
		"firstname":  firstname,
		"password":   password,
		"role":       role,
		"created_at": now,
		"updated_at": now,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(username, lastname, firstname, password, role, manager string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()

	_, err = userCollection.InsertOne(ctx, bson.M{
		"username":   username,
		"lastname":   lastname,
		"firstname":  firstname,
		"password":   string(hashedPassword),
		"role":       role,
		"manager":    manager,
		"created_at": now,
		"updated_at": now,
	})

	return err
}

// GetUser retrieves a user by username from db
func GetUserByUsername(username string) (models.User, error) {
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

// UpdateUser updates user information in the database
func UpdateUser(username, lastname, firstname, password, role string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()

	_, err := userCollection.UpdateOne(ctx, bson.M{"username": username}, bson.M{
		"$set": bson.M{
			"lastname":   lastname,
			"firstname":  firstname,
			"password":   password,
			"role":       role,
			"updated_at": now,
		},
	})

	return err
}

// DeleteUser deletes a user from the database
func DeleteUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userCollection.DeleteOne(ctx, bson.M{"username": username})
	return err
}

// ListAllUsers retrieves all users from the database
func ListAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// ListUsersByManager retrieves users under a specific manager
func ListUsersByManager(managerUsername string) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := userCollection.Find(ctx, bson.M{"manager": managerUsername})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
