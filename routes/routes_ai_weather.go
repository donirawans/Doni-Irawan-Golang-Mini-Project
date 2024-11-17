package routes

import (
	"warningfloodsystem/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterAIWeatherRoutes(e *echo.Echo, ctrl *controllers.AIWeatherController) {
	e.GET("/weather/ai/:id", ctrl.GetRecommendation) 
}
