package main

import (
	"log"
	"warningfloodsystem/config"
	"warningfloodsystem/controllers"
	"warningfloodsystem/domain/repository"
	"warningfloodsystem/domain/usecase"
	"warningfloodsystem/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Authentication setup
	authRepo := repository.NewAuthRepository(db)
	authService := usecase.NewAuthUsecase(authRepo)
	authController := controllers.NewAuthController(authService)

	// Flood report setup
	floodReportRepo := repository.NewFloodReportRepository(db)
	floodReportService := usecase.NewFloodReportUsecase(*floodReportRepo)
	floodReportController := controllers.NewFloodReportController(floodReportService)

	// Weather data setup
	weatherDataRepo := repository.NewWeatherDataRepository(db)
	WeatherDataUsecase := usecase.NewWeatherDataService(weatherDataRepo)
	weatherDataController := controllers.NewWeatherDataController(WeatherDataUsecase)

	// AI recommendation setup
	AiRecommendationRepository := repository.NewAIWeatherRepository()
	AiRecommendationUsecase := usecase.NewAIWeatherService(AiRecommendationRepository)
	AiRecommendationController := controllers.NewAIWeatherController(AiRecommendationUsecase, WeatherDataUsecase)

	// Initialize Echo
	e := echo.New()
	e.Use(middleware.Logger())

	// Register routes
	routes.RegisterRoutes(e, authController)
	routes.RegisterReportRoutes(e, floodReportController)
	routes.RegisterWeatherDataRoutes(e, weatherDataController)
	routes.RegisterAIWeatherRoutes(e, AiRecommendationController)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
