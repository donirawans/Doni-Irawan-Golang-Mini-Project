package model

import "time"

type FloodReport struct {
	ID           uint      `json:"id" gorm:"primaryKey;column:data_id"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	Lokasi       string    `json:"lokasi" gorm:"not null"`
	WaktuLaporan time.Time `json:"waktu_laporan" gorm:"not null"`
	Deskripsi    string    `json:"deskripsi" gorm:"type:text"`
	Foto         string    `json:"foto" gorm:"size:255"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}
