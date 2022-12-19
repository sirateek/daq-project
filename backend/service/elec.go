package service

import (
	"happyEWater/models"
	"happyEWater/repository"
	"time"
)

type electricService struct {
	electricRepository repository.ElectricRepocitory
}

type ElectricService interface {
	GetElectric() ([]models.Electric, error)
	GetDate() ([]time.Time, error)
	GetElecPerDay() ([]float32, error)
	GetElec3Hour() ([]models.ElectricTmd, error)
}

func NewElectricService(electricRepository repository.ElectricRepocitory) ElectricService {
	return &electricService{
		electricRepository: electricRepository,
	}
}

func (e *electricService) GetElectric() ([]models.Electric, error) {
	data, err := e.electricRepository.GetElectric()
	return data, err
}

func (e *electricService) GetDate() ([]time.Time, error) {
	data, err := e.electricRepository.GetDate()
	return data, err
}

func (e *electricService) GetElecPerDay() ([]float32, error) {
	data, err := e.electricRepository.GetElecPerDay()
	return data, err
}

func (e *electricService) GetElec3Hour() ([]models.ElectricTmd, error) {
	data, err := e.electricRepository.GetElec3Hour()
	return data, err
}
