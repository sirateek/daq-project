package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TMD struct {
	ID              primitive.ObjectID `bson:"_id"`
	WeatherForcasts []interface{}      `bson:"WeatherForecasts, omitempty"`
	Temp            float64            `bson:"temp, omitempty"`
	Rainfall        int                `bson:"rain"`
	TimeStamp       int64              `bson:"timestamp"`
}

type RawTMD struct {
	Temp      float64   `json:"temp"`
	Rainfall  int       `json:"rain"`
	TimeStamp time.Time `json:"timestamp"`
}
