package repository

import (
	"context"
	"happyEWater/models"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type tmdRepository struct {
	tmdCollection *mongo.Collection
}

type TmdRepository interface {
	GetTmd() ([]models.TMD, error)
	GetTmdPerDay() ([]float64, []int, error)
	GetDate() ([]time.Time, error)
	GetAllDate() ([]time.Time, error)
	GetRawData() ([]models.RawTMD, error)
}

func NewTmdRepository(tmdCollection *mongo.Collection) TmdRepository {
	return &tmdRepository{
		tmdCollection: tmdCollection,
	}
}

func (t *tmdRepository) GetTmd() ([]models.TMD, error) {
	var listTmdData []models.TMD
	var tmdData models.TMD
	data, err := t.tmdCollection.Find(context.Background(), bson.M{})
	if err != nil {
		logrus.Info(err)
		return []models.TMD{}, err
	}
	defer data.Close(context.Background())
	for data.Next(context.Background()) {
		err := data.Decode(&tmdData)
		if err != nil {
			logrus.Info(err)
		}
		listTmdData = append(listTmdData, tmdData)
	}
	return listTmdData, nil
}

func (t *tmdRepository) GetTmdPerDay() ([]float64, []int, error) {
	var tmdData models.TMD
	var listTemp []float64
	var listRain []int
	var baseDate time.Time
	var currentTemp float64
	var currentRain int
	var count int = 1

	data, err := t.tmdCollection.Find(context.Background(), bson.M{})
	if err != nil {
		logrus.Info(err)
		return []float64{}, []int{}, err
	}
	defer data.Close(context.Background())
	for data.Next(context.Background()) {
		data.Decode(&tmdData)
		dataDate := ConvertTime(int64(tmdData.TimeStamp)).UTC()
		if baseDate.Day() != dataDate.Day() || baseDate.Month() != dataDate.Month() || baseDate.Year() != dataDate.Year() {
			baseDate = dataDate
			listTemp = append(listTemp, currentTemp/float64(count))
			listRain = append(listRain, currentRain/count)
			currentTemp = 0
			currentRain = 0
			count = 1
		}
		currentRain += tmdData.Rainfall
		currentTemp += tmdData.Temp
		count += 1
	}
	return listTemp, listRain, nil
}

func (t *tmdRepository) GetDate() ([]time.Time, error) {
	var tmdData models.TMD
	var date []time.Time

	data, err := t.tmdCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return []time.Time{}, ErrNotFound
	}
	for data.Next(context.Background()) {
		data.Decode(&tmdData)
		dataDate := ConvertTime(int64(tmdData.TimeStamp)).UTC()
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

func (t *tmdRepository) GetAllDate() ([]time.Time, error) {
	var tmdData models.TMD
	var date []time.Time

	data, err := t.tmdCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return []time.Time{}, ErrNotFound
	}
	for data.Next(context.Background()) {
		data.Decode(&tmdData)
		dataDate := ConvertTime(int64(tmdData.TimeStamp)).UTC()
		date = append(date, dataDate)
	}
	return date, nil
}

func (t *tmdRepository) GetRawData() ([]models.RawTMD, error) {
	var listTmdData []models.RawTMD
	var tmdData models.TMD
	data, err := t.tmdCollection.Find(context.Background(), bson.M{})
	if err != nil {
		logrus.Info(err)
		return []models.RawTMD{}, err
	}
	defer data.Close(context.Background())
	for data.Next(context.Background()) {
		err := data.Decode(&tmdData)
		if err != nil {
			logrus.Info(err)
		}
		listTmdData = append(listTmdData, models.RawTMD{Temp: tmdData.Temp, Rainfall: tmdData.Rainfall, TimeStamp: ConvertTime(int64(tmdData.TimeStamp)).UTC()})
	}
	return listTmdData, nil
}
