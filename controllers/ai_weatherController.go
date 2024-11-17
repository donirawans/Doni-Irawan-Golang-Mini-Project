package controllers

import (
	"log"
	"net/http"
	"strconv"
	"warningfloodsystem/domain/usecase"
	"warningfloodsystem/utils"

	"github.com/labstack/echo/v4"
)

type AIWeatherController struct {
	aiUsecase   usecase.AIWeatherUsecase
	dataUsecase usecase.WeatherDataUsecase
}

func NewAIWeatherController(aiUsecase usecase.AIWeatherUsecase, dataUsecase usecase.WeatherDataUsecase) *AIWeatherController {
	return &AIWeatherController{
		aiUsecase:   aiUsecase,
		dataUsecase: dataUsecase,
	}
}

type WeatherResponse struct {
	WeatherData         interface{} `json:"weather_data"`
	PrediksiRisikoBanjir struct {
		Rekomendasi string `json:"rekomendasi"`
		PredictedAt string `json:"predicted_at"`
	} `json:"Prediksi_Risiko_Banjir"`
}

func (ctrl *AIWeatherController) GetRecommendation(ctx echo.Context) error {
	// Ambil ID dari path parameter
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("[ERROR] Invalid ID format:", err)
		return utils.JSONErrorResponse(ctx, http.StatusBadRequest, "Invalid ID format")
	}

	// Ambil data cuaca berdasarkan ID
	weatherData, err := ctrl.dataUsecase.GetByID(uint(id))
	if err != nil {
		log.Println("[ERROR] Weather data not found:", err)
		return utils.JSONErrorResponse(ctx, http.StatusNotFound, "Weather data not found")
	}

	// Validasi apakah data cuaca valid untuk AI
	if weatherData.CurahHujan <= 0 || weatherData.TinggiSungai <= 0 {
		return utils.JSONErrorResponse(ctx, http.StatusBadRequest, "Invalid weather data for recommendation")
	}

	// Panggil AI untuk rekomendasi
	recommendation, err := ctrl.aiUsecase.GetRecommendation(ctx.Request().Context(), weatherData.CurahHujan, weatherData.TinggiSungai)
	if err != nil {
		log.Println("[ERROR] Failed to get AI recommendation:", err)
		return utils.JSONErrorResponse(ctx, http.StatusInternalServerError, "Failed to get AI recommendation")
	}

	// Buat respons JSON dengan struktur yang benar
	response := WeatherResponse{
		WeatherData: weatherData,
	}
	response.PrediksiRisikoBanjir.Rekomendasi = recommendation.Rekomendasi
	response.PrediksiRisikoBanjir.PredictedAt = recommendation.PredictedAt.Format("2006-01-02T15:04:05Z07:00")

	return utils.JSONSuccessResponse(ctx, response)
}
