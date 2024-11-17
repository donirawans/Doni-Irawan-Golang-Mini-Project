package controllers

import (
	"net/http"
	"warningfloodsystem/domain/model"
	"warningfloodsystem/domain/usecase"
	"warningfloodsystem/utils"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase usecase.AuthUsecase
}

func NewAuthController(u usecase.AuthUsecase) *AuthController {
	return &AuthController{authUseCase: u}
}

func (a *AuthController) Register(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := a.authUseCase.RegisterUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to register"})
	}

	return utils.JSONSuccessResponse(c, "Registration successful")
}

func (a *AuthController) Login(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	token, err := a.authUseCase.Login(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	response := map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	}
	return utils.JSONSuccessResponse(c, response)
}
