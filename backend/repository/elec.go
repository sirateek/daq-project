package repository

import (
	"context"
	"fmt"
	"happyEWater/models"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type electricRepocitory struct {
	electricCollection *mongo.Collection
}

type ElectricRepocitory interface {
	GetElectric() ([]models.Electric, error)
	GetDate() ([]time.Time, error)
	GetElecPerDay() ([]float32, error)
	GetElec3Hour() ([]models.ElectricTmd, error)
}

func NewElectricRepository(electricCollection *mongo.Collection) ElectricRepocitory {
	return &electricRepocitory{
		electricCollection: electricCollection,
	}
}

func (e *electricRepocitory) GetElectric() ([]models.Electric, error) {
	var electricData models.Electric
	var listElectricData []models.Electric
	data, err := e.electricCollection.Find(context.Background(), bson.M{})
	if err != nil {
		logrus.Fatal(err)
		return []models.Electric{}, err
	}
	defer data.Close(context.Background())
	for data.Next(context.Background()) {
		data.Decode(&electricData)
		listElectricData = append(listElectricData, electricData)
	}
	return listElectricData, err
}

func (e *electricRepocitory) GetDate() ([]time.Time, error) {
	var elecData models.Electric
	var date []time.Time

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", 1}})
	data, err := e.electricCollection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		return []time.Time{}, ErrNotFound
	}
	for data.Next(context.Background()) {
		data.Decode(&elecData)
		dataDate := ConvertTime(int64(elecData.Timestamp)).UTC()
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

func (e *electricRepocitory) GetElecPerDay() ([]float32, error) {
	var elecData models.Electric
	var listElec []float32
	var baseDate time.Time
	var currentElec float32
	var count float32 = 1
	ctx := context.Background()
	data, err := e.electricCollection.Find(ctx, bson.D{})
	if err != nil {
		logrus.Info(err)
		return []float32{}, err
	}
	for data.Next(ctx) {
		data.Decode(&elecData)
		dataDate := ConvertTime(int64(elecData.Timestamp)).UTC()
		if baseDate.Day() != dataDate.Day() || baseDate.Month() != dataDate.Month() || baseDate.Year() != dataDate.Year() {
			baseDate = dataDate
			listElec = append(listElec, currentElec/count)
			currentElec = 0
			count = 1
		}
		currentElec += float32(elecData.Watt)
		count += 1
	}
	listElec = append(listElec, currentElec)
	data.Close(ctx)
	return listElec, nil
}

func (e *electricRepocitory) GetElec3Hour() ([]models.ElectricTmd, error) {
	var listElec []models.ElectricTmd
	var baseDate time.Time
	var currentElec float64 = 0
	ctx := context.Background()
	data, err := e.electricCollection.Find(ctx, bson.D{})
	if err != nil {
		logrus.Info(err)
		return []models.ElectricTmd{}, err
	}
	// size := 0
	for data.Next(ctx) {
		var elecData models.Electric
		err := data.Decode(&elecData)
		if err != nil {
			logrus.Error(err)
			continue
		}
		fmt.Println(elecData)
		dataDate := ConvertTime(int64(elecData.Timestamp)).UTC()
		if baseDate.IsZero() {
			baseDate = dataDate.Add(time.Hour * 3)
		}
		if dataDate.After(baseDate) {
			listElec = append(listElec, models.ElectricTmd{Watt: currentElec / 3600000, TimeStamp: baseDate})
			baseDate = baseDate.Add(time.Hour * 3)
			currentElec = 0
		}
		currentElec += elecData.Watt
		// print(currentElec)
		// size++
	}
	// fmt.Println("Size: ", size)
	listElec = append(listElec, models.ElectricTmd{Watt: currentElec / 3600000, TimeStamp: baseDate})
	data.Close(ctx)
	return listElec, nil
}
