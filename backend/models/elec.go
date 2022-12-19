package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Electric struct {
	ID        primitive.ObjectID `bson:"_id"`
	Amp       float64            `bson:"amp, omitempty"`
	Watt      float64            `bson:"watt, omitemepty"`
	Volt      int                `bson:"volt, omitemepty"`
	Timestamp int64              `bson:"timestamp"`
}

type ElectricTmd struct {
	Watt      float64   `json:"electric" bson:"watt"`
	TimeStamp time.Time `json:"timestamp" bson:"timestamp"`
}
