package repository

import (
	"fmt"
	"warningfloodsystem/domain/model"

	"gorm.io/gorm"
)

type WeatherDataRepository interface {
	Create(weatherData *model.WeatherData) (*model.WeatherData, error)
	GetByID(weatherDataID uint) (*model.WeatherData, error)
	GetAll() ([]model.WeatherData, error)
	Update(weatherData *model.WeatherData) (*model.WeatherData, error)
	Delete(id uint) error
}

type WeatherDataRepo struct {
	db *gorm.DB
}

func NewWeatherDataRepository(db *gorm.DB) *WeatherDataRepo {
	return &WeatherDataRepo{db: db}
}

func (r *WeatherDataRepo) Create(weatherData *model.WeatherData) (*model.WeatherData, error) {
	if err := r.db.Create(weatherData).Error; err != nil {
		return nil, fmt.Errorf("terjadi kesalahan saat membuat data cuaca: %w", err)
	}
	return weatherData, nil
}

func (r *WeatherDataRepo) GetByID(weatherDataID uint) (*model.WeatherData, error) {
	var weatherData model.WeatherData
	err := r.db.First(&weatherData, weatherDataID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("data cuaca dengan ID %d tidak ditemukan", weatherDataID)
		}
		// Error lainnya
		return nil, fmt.Errorf("terjadi kesalahan saat mengambil data cuaca dengan ID %d: %w", weatherDataID, err)
	}
	return &weatherData, nil
}

func (r *WeatherDataRepo) GetAll() ([]model.WeatherData, error) {
	var weatherData []model.WeatherData
	err := r.db.Find(&weatherData).Error
	if err != nil {
		return nil, fmt.Errorf("terjadi kesalahan saat mengambil semua data cuaca: %w", err)
	}
	return weatherData, nil
}

func (r *WeatherDataRepo) Update(weatherData *model.WeatherData) (*model.WeatherData, error) {
	if err := r.db.Model(&model.WeatherData{}).Where("id = ?", weatherData.ID).Updates(weatherData).Error; err != nil {
		return nil, fmt.Errorf("terjadi kesalahan saat memperbarui data cuaca dengan ID %d: %w", weatherData.ID, err)
	}
	return weatherData, nil
}


func (r *WeatherDataRepo) Delete(id uint) error {
	if err := r.db.Delete(&model.WeatherData{}, id).Error; err != nil {
		return fmt.Errorf("terjadi kesalahan saat menghapus data cuaca dengan ID %d: %w", id, err)
	}
	return nil
}
