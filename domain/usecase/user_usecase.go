package usecase

import (
	"errors"
	"warningfloodsystem/domain/model"
	"warningfloodsystem/domain/repository"
	"warningfloodsystem/middlewares"
	"warningfloodsystem/security"
)

type AuthUsecase interface {
	RegisterUser(user *model.User) error
	Login(email, password string) (string, error)
}

type AuthUsecaseImpl struct {
	repo repository.UserRepository
}

func NewAuthUsecase(repo repository.UserRepository) AuthUsecase {
	return &AuthUsecaseImpl{repo}
}

func (s *AuthUsecaseImpl) RegisterUser(user *model.User) error {
	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}

func (s *AuthUsecaseImpl) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if !security.CheckPasswordHash(password, user.Password) {
		return "", errors.New("kredensial tidak valid")
	}
	return middlewares.GenerateJWT(uint(user.ID))
}

func Add(a, b int) int {
	result := a + b
	if result < 0 {
		return 0
	}
	return result
}
