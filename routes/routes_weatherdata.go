package routes

import (
	"warningfloodsystem/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterWeatherDataRoutes(e *echo.Echo, ctrl *controllers.WeatherDataController) {
	e.POST("/weather", ctrl.Create)
	e.GET("/weather", ctrl.GetAll)
	e.GET("/weather/:id", ctrl.GetByID)
	e.PUT("/weather/:id", ctrl.Update)
	e.DELETE("/weather/:id", ctrl.Delete)
}
