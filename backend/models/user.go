package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Role      string             `json:"role" bson:"role"`
	Manager   string             `json:"manager,omitempty" bson:"manager,omitempty"` // Added for manager relationship
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

const (
	AdminRole   = "admin"
	ManagerRole = "manager"
	UserRole    = "user"
)
