package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Water struct {
	ID        primitive.ObjectID `bson:"_id"`
	FlowRate  float64            `bson:"flowRate"`
	TimeStamp int                `bson:"timeStamp"`
	Water     float64            `bson:"water"`
}

type WaterTmd struct {
	Water     float32   `json:"electric" bson:"watt"`
	TimeStamp time.Time `json:"timestamp" bson:"timestamp"`
}
