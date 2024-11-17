package model

import "time"

type AIWeather struct {
	ID          uint      			`json:"id" gorm:"primaryKey;column:ai_weather_id"`
	Rekomendasi string    			`json:"rekomendasi" gorm:"type:text"`
	PredictedAt time.Time 			`json:"predicted_at"`                
}
