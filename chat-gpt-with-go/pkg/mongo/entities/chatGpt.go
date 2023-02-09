package entities

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatGpt struct {
    ID                  primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
    Prompt              string `json:"prompt" bson:"prompt"`
    Output              string `json:"output" bson:"output"`
    CreatedAt           int64  `json:"created_at" bson:"created_at"`
    CreatedBy           string  `json:"created_by" bson:"created_by"`
}
