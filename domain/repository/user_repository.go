package repository

import (
	"fmt"
	"warningfloodsystem/domain/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUser(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("terjadi kesalahan saat membuat pengguna: %w", err)
	}
	return nil
}

func (r *AuthRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("pengguna dengan email %s tidak ditemukan", email)
		}
		return nil, fmt.Errorf("terjadi kesalahan saat mencari pengguna dengan email %s: %w", email, err)
	}
	return &user, nil
}
