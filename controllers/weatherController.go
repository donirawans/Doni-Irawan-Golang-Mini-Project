package controllers

import (
	"log"
	"net/http"
	"strconv"
	"warningfloodsystem/domain/model"
	"warningfloodsystem/domain/usecase"
	"warningfloodsystem/utils"

	"github.com/labstack/echo/v4"
)

type WeatherDataController struct {
	service usecase.WeatherDataUsecase
}

func NewWeatherDataController(service usecase.WeatherDataUsecase) *WeatherDataController {
	return &WeatherDataController{service: service}
}

func (ctrl *WeatherDataController) Create(ctx echo.Context) error {
	var weatherData model.WeatherData
	if err := ctx.Bind(&weatherData); err != nil {
		log.Println("[ERROR] Failed to bind data:", err)
		return utils.JSONErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	createdData, err := ctrl.service.Create(&weatherData)
	if err != nil {
		log.Println("[ERROR] Failed to create weather data:", err)
		return utils.JSONErrorResponse(ctx, http.StatusInternalServerError, "Failed to create weather data")
	}

	return utils.JSONSuccessResponse(ctx, createdData)
}

func (ctrl *WeatherDataController) GetAll(ctx echo.Context) error {
	data, err := ctrl.service.GetAll()
	if err != nil {
		log.Println("[ERROR] Failed to fetch weather data:", err)
		return utils.JSONErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch weather data")
	}

	return utils.JSONSuccessResponse(ctx, data)
}

func (ctrl *WeatherDataController) GetByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("[ERROR] Invalid ID format:", err)
		return utils.JSONErrorResponse(ctx, http.StatusBadRequest, "Invalid ID format")
	}

	data, err := ctrl.service.GetByID(uint(id))
	if err != nil {
		log.Println("[ERROR] Weather data not found:", err)
		return utils.JSONErrorResponse(ctx, http.StatusNotFound, "Weather data not found")
	}

	return utils.JSONSuccessResponse(ctx, data)
}

func (c *WeatherDataController) Update(ctx echo.Context) error {
	var weatherData model.WeatherData

	if err := ctx.Bind(&weatherData); err != nil {
		log.Printf("[ERROR] Failed to bind data for update: %v", err)
		return utils.JSONErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("[ERROR] Invalid ID format: %v", err)
		return utils.JSONErrorResponse(ctx, http.StatusBadRequest, "Invalid ID format")
	}
	weatherData.ID = uint(id)

	updatedWeatherData, err := c.service.Update(&weatherData)
	if err != nil {
		log.Printf("[ERROR] Failed to update weather data (ID: %d): %v", id, err)
		return utils.JSONErrorResponse(ctx, http.StatusInternalServerError, "Failed to update weather data")
	}

	return utils.JSONSuccessResponse(ctx, updatedWeatherData)
}

func (c *WeatherDataController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("[ERROR] Invalid ID format: %v", err)
		return utils.JSONErrorResponse(ctx, http.StatusBadRequest, "Invalid ID format")
	}

	err = c.service.Delete(uint(id))
	if err != nil {
		log.Printf("[ERROR] Failed to delete weather data (ID: %d): %v", id, err)
		return utils.JSONErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete weather data")
	}

	return utils.JSONSuccessResponse(ctx, "Weather data deleted successfully")
}
