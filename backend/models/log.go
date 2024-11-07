package models

type Log struct {
	ID		string `json:"id" bson:"_id,omitempty"`
	ScriptID	string `json:"script_id" bson:"script_id"`
	Message		string `json:"message" bson:"message"`
	Timestamp	int64 `json:"timestamp" bson:"timestamp"`
}