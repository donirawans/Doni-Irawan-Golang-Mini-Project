package routes

import (
	"warningfloodsystem/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterReportRoutes(e *echo.Echo, ctrl *controllers.FloodReportController) {
	e.POST("/floodreports", ctrl.Create)
	e.GET("/floodreports", ctrl.GetAll)
	e.GET("/floodreports/:id", ctrl.GetByID)
	e.PUT("/floodreports/:id", ctrl.Update)
	e.DELETE("/floodreports/:id", ctrl.Delete)
}