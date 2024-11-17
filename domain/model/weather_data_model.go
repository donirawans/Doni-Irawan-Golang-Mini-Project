package model

import "time"

type WeatherData struct {
	ID           uint      `json:"id" gorm:"primaryKey;column:weather_id"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	CurahHujan   float64   `json:"curah_hujan" gorm:"not null"`
	TinggiSungai float64   `json:"tinggi_sungai" gorm:"not null"`
	RecordedAt   time.Time `json:"recorded_at" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}
