package service

import (
	"happyEWater/models"
	"happyEWater/repository"
	"time"
)

type tmdService struct {
	tmdRepository repository.TmdRepository
}

type TmdService interface {
	GetTmd() ([]models.TMD, error)
	GetTimePerDay() ([]float64, []int, error)
	GetDate() ([]time.Time, error)
	GetAllDate() ([]time.Time, error)
	GetRawData() ([]models.RawTMD, error)
}

func NewTmdService(tmdRepository repository.TmdRepository) TmdService {
	return &tmdService{
		tmdRepository: tmdRepository,
	}
}

func (t *tmdService) GetTmd() ([]models.TMD, error) {
	data, err := t.tmdRepository.GetTmd()
	return data, err
}

func (t *tmdService) GetTimePerDay() ([]float64, []int, error) {
	temp, rain, err := t.tmdRepository.GetTmdPerDay()
	return temp, rain, err
}

func (t *tmdService) GetDate() ([]time.Time, error) {
	data, err := t.tmdRepository.GetDate()
	return data, err
}

func (t *tmdService) GetAllDate() ([]time.Time, error) {
	data, err := t.tmdRepository.GetAllDate()
	return data, err
}

func (t *tmdService) GetRawData() ([]models.RawTMD, error) {
	data, err := t.tmdRepository.GetRawData()
	return data, err
}
