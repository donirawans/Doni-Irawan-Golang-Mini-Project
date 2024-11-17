package routes

import (
	"warningfloodsystem/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, authController *controllers.AuthController) {
	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)
}