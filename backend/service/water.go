package service

import (
	"happyEWater/models"
	"happyEWater/repository"
	"time"
)

type waterService struct {
	waterRepository repository.WaterRepository
}

type WaterService interface {
	GetWater() ([]models.Water, error)
	GetWaterPerDay() ([]float64, error)
	GetDate() ([]time.Time, error)
	GetWater3Hour() ([]models.WaterTmd, error)
}

func NewWaterService(waterRepocitory repository.WaterRepository) WaterService {
	return &waterService{
		waterRepository: waterRepocitory,
	}
}

func (w *waterService) GetWater() ([]models.Water, error) {
	data, err := w.waterRepository.GetWater()
	return data, err
}

func (w *waterService) GetWaterPerDay() ([]float64, error) {
	data, err := w.waterRepository.GetWaterPerDay()
	return data, err
}

func (w *waterService) GetDate() ([]time.Time, error) {
	data, err := w.waterRepository.GetDate()
	return data, err
}

func (w *waterService) GetWater3Hour() ([]models.WaterTmd, error) {
	data, err := w.waterRepository.GetWater3Hour()
	return data, err
}
