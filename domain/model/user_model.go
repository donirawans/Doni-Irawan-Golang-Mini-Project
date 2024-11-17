package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;column:user_id"`
	Name      string    `json:"username" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
