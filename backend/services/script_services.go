package services

import (
	"bytes"
	"context"
	"os/exec"
	"time"

	"github.com/dokuqui/monitor_scheduler/backend/db"
	"github.com/dokuqui/monitor_scheduler/backend/models"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var scriptCollection *mongo.Collection
var chronScheduler *cron.Cron

func InitializeScriptCollection() {
	scriptCollection = db.Client.Database("monitorScheduler").Collection("scripts")
	chronScheduler = cron.New()
	chronScheduler.Start()
}

func GetScriptByID(scriptID string) (models.Script, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var script models.Script

	err := scriptCollection.FindOne(ctx, bson.M{"_id": scriptID}).Decode(&script)

	return script, err
}

func GetScriptsByUser(username string) ([]models.Script, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var scripts []models.Script

	cursor, err := scriptCollection.Find(ctx, bson.M{"owner": username})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &scripts); err != nil {
		return nil, err
	}

	return scripts, nil
}

func GetScriptsByUserGroup(userGroup string) ([]models.Script, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var scripts []models.Script

	cursor, err := scriptCollection.Find(ctx, bson.M{"user_group": userGroup})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &scripts); err != nil {
		return nil, err
	}

	return scripts, nil
}

func GetAllScripts() ([]models.Script, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var scripts []models.Script

	cursor, err := scriptCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &scripts); err != nil {
		return nil, err
	}

	return scripts, nil
}

func CreateScript(name, content, owner, userGroup string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	script := models.Script{
		ID:        primitive.NewObjectID().Hex(),
		Name:      name,
		Content:   content,
		Owner:     owner,
		UserGroup: userGroup,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := scriptCollection.InsertOne(ctx, script)
	return err
}

func UpdateScript(script models.Script) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	script.UpdatedAt = time.Now()

	_, err := scriptCollection.UpdateOne(ctx, bson.M{"_id": script.ID}, bson.M{
		"$set": script,
	})

	return err
}

func DeleteScript(scriptID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := scriptCollection.DeleteOne(ctx, bson.M{"_id": scriptID})

	return err
}

func UploadScript(script models.Script) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	script.ID = primitive.NewObjectID().Hex()
	script.CreatedAt = time.Now()
	script.UpdatedAt = time.Now()

	_, err := scriptCollection.InsertOne(ctx, script)

	return err
}

func ExecuteScript(scriptID string) (string, error) {
	script, err := GetScriptByID(scriptID)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("bash", "-c", script.Content)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()
	if err != nil {
		LogMessage(scriptID, err.Error())
		return "", err
	}

	// Log the output of the script
	err = LogMessage(scriptID, out.String())
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func ScheduleScript(scriptID string, scheduleTime time.Time) error {
	// Calculate the delay duration
	cronExpr := scheduleTime.Format("05 04 15 02 01 ?")

	// Schedule the script execution
	_, err := chronScheduler.AddFunc(cronExpr, func() {
		_, errcronExpr := ExecuteScript(scriptID)
		if errcronExpr != nil {
			LogMessage(scriptID, errcronExpr.Error())
		}
	})

	return err
}
