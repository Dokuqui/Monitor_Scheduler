package models

import "time"

type Script struct {
	ID string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Content string `json:"content" bson:"content"`
	Owner string `json:"owner" bson:"owner"`
	UserGroup string `json:"user_group" bson:"user_group"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}