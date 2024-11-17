package usecase

import (
	"errors"
	"warningfloodsystem/domain/model"
	"warningfloodsystem/domain/repository"
)

type WeatherDataUsecase interface {
	Create(weatherData *model.WeatherData) (*model.WeatherData, error)
	GetByID(weatherDataID uint) (*model.WeatherData, error)
	GetAll() ([]model.WeatherData, error)
	Update(weatherData *model.WeatherData) (*model.WeatherData, error)
	Delete(weatherDataID uint) error
}

type WeatherDataService struct {
	repo repository.WeatherDataRepository
}

func NewWeatherDataService(r repository.WeatherDataRepository) *WeatherDataService {
	return &WeatherDataService{
		repo: r,
	}
}

func (s *WeatherDataService) Create(weatherData *model.WeatherData) (*model.WeatherData, error) {
	if weatherData.CurahHujan <= 0 || weatherData.TinggiSungai <= 0 {
		return nil, errors.New("curah hujan dan tinggi sungai harus lebih besar dari 0")
	}
	return s.repo.Create(weatherData)
}

func (s *WeatherDataService) GetAll() ([]model.WeatherData, error) {
	return s.repo.GetAll()
}

func (s *WeatherDataService) GetByID(id uint) (*model.WeatherData, error) {
	return s.repo.GetByID(id)
}

func (s *WeatherDataService) Update(weatherData *model.WeatherData) (*model.WeatherData, error) {
	if weatherData.ID == 0 {
		return nil, errors.New("ID diperlukan untuk memperbarui data")
	}
	return s.repo.Update(weatherData)
}

func (s *WeatherDataService) Delete(id uint) error {
	return s.repo.Delete(id)
}
