package services

import (
	"context"
	"time"

	"github.com/dokuqui/monitor_scheduler/backend/db"
	"github.com/dokuqui/monitor_scheduler/backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var logCollection *mongo.Collection

func InitializeLogCollection() {
    logCollection = db.Client.Database("monitorScheduler").Collection("logs")
}

func LogMessage(scriptID string, message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log := models.Log{
		ScriptID:  scriptID,
		Message:   message,
		Timestamp: time.Now().Unix(),
	}

	_, err := logCollection.InsertOne(ctx, log)
	return err
}