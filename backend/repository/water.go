package repository

import (
	"context"
	"errors"
	"happyEWater/models"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type waterRepository struct {
	waterCollection *mongo.Collection
}

type WaterRepository interface {
	GetWater() ([]models.Water, error)
	GetWaterPerDay() ([]float64, error)
	GetDate() ([]time.Time, error)
	GetWater3Hour() ([]models.WaterTmd, error)
}

func NewWaterCollection(waterCollection *mongo.Collection) WaterRepository {
	return &waterRepository{
		waterCollection: waterCollection,
	}
}

func (w *waterRepository) GetWater() ([]models.Water, error) {
	var waterListData []models.Water
	var waterData models.Water
	matchStage := bson.D{{"$match", bson.D{{"flowRate", bson.D{{"$gt", 0}, {"$lt", 100}}}}}}
	data, _ := w.waterCollection.Aggregate(context.TODO(), mongo.Pipeline{matchStage})
	if data == nil {
		return []models.Water{}, ErrNotFound
	}
	for data.Next(context.Background()) {
		errs := data.Decode(&waterData)
		if errs != nil {
			logrus.Info(errs)
		}
		waterListData = append(waterListData, waterData)
	}
	return waterListData, nil
}

func (w *waterRepository) GetWaterPerDay() ([]float64, error) {
	var waterData models.Water
	var listWater []float64
	var baseDate time.Time
	var count float64 = 1
	var currentWater float64

	matchStage := bson.D{{"$match", bson.D{{"flowRate", bson.D{{"$gt", 0}, {"$lt", 100}}}}}}
	data, err := w.waterCollection.Aggregate(context.Background(), mongo.Pipeline{matchStage})
	if err != nil {
		logrus.Info(err)
		return []float64{}, err
	}
	defer data.Close(context.Background())
	for data.Next(context.Background()) {
		data.Decode(&waterData)
		dataDate := ConvertTime(int64(waterData.TimeStamp)).UTC()
		if baseDate.Day() != dataDate.Day() || baseDate.Month() != dataDate.Month() || baseDate.Year() != dataDate.Year() {
			baseDate = dataDate
			listWater = append(listWater, currentWater/count)
			currentWater = 0
			count = 1
		}
		currentWater += waterData.Water
		count += 1
	}
	listWater = append(listWater, currentWater)
	return listWater, nil
}

func ConvertTime(milli int64) time.Time {
	return time.Unix(0, milli*int64(time.Millisecond))
}

var ErrNotFound error = errors.New("Some thing happen.")

func (w *waterRepository) GetDate() ([]time.Time, error) {
	var waterData models.Water
	var date []time.Time

	matchStage := bson.D{{"$match", bson.D{{"flowRate", bson.D{{"$gt", 0}, {"$lt", 100}}}}}}
	data, _ := w.waterCollection.Aggregate(context.TODO(), mongo.Pipeline{matchStage})
	if data == nil {
		return []time.Time{}, ErrNotFound
	}

	for data.Next(context.Background()) {
		data.Decode(&waterData)
		dataDate := ConvertTime(int64(waterData.TimeStamp)).UTC()
		if date != nil {
			closeDate := date[len(date)-1]
			if closeDate.Day() == dataDate.Day() && closeDate.Month() == dataDate.Month() && closeDate.Year() == dataDate.Year() {
				continue
			}
		}
		date = append(date, dataDate)
	}
	return date, nil
}

func (w *waterRepository) GetWater3Hour() ([]models.WaterTmd, error) {
	var waterData models.Water
	var listWater []models.WaterTmd
	var baseDate time.Time
	var currentWater float32
	var count float32 = 1

	matchStage := bson.D{{"$match", bson.D{{"flowRate", bson.D{{"$gt", 0}, {"$lt", 100}}}}}}
	data, _ := w.waterCollection.Aggregate(context.TODO(), mongo.Pipeline{matchStage})
	if data == nil {
		return []models.WaterTmd{}, ErrNotFound
	}
	defer data.Close(context.Background())
	for data.Next(context.Background()) {
		data.Decode(&waterData)
		dataDate := ConvertTime(int64(waterData.TimeStamp)).UTC()
		if baseDate.IsZero() {
			baseDate = dataDate.Add(time.Hour * 3)
		}
		if dataDate.After(baseDate) {
			listWater = append(listWater, models.WaterTmd{Water: currentWater, TimeStamp: baseDate})
			baseDate = baseDate.Add(time.Hour * 3)
			currentWater = 0
			count = 1
		}
		currentWater += float32(waterData.Water)
		count += 1
	}
	listWater = append(listWater, models.WaterTmd{Water: currentWater, TimeStamp: baseDate})
	return listWater, nil
}
