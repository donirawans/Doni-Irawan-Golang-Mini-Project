package usecase

import (
	"context"
	"fmt"
	"warningfloodsystem/domain/model"
	"warningfloodsystem/domain/repository"
)

type AIWeatherUsecase interface {
	GetRecommendation(ctx context.Context, curahHujan, tinggiSungai float64) (*model.AIWeather, error)
}

type aiWeatherService struct {
	repo repository.AIWeatherRepository
}

func NewAIWeatherService(repo repository.AIWeatherRepository) AIWeatherUsecase {
	return &aiWeatherService{repo: repo}
}

func (s *aiWeatherService) GetRecommendation(ctx context.Context, curahHujan, tinggiSungai float64) (*model.AIWeather, error) {
	question := fmt.Sprintf(
		"Dengan curah hujan %.2f mm dan tinggi sungai %.2f mm, apa risiko banjirnya? Jelaskan mitigasi yang dapat diambil untuk mengurangi dampak banjir.",
		curahHujan, tinggiSungai,
	)
	response, err := s.repo.GetAIResponse(ctx, question, curahHujan, tinggiSungai)
	if err != nil {
		return nil, err
	}

	return &model.AIWeather{
		Rekomendasi: response.Rekomendasi,
		PredictedAt: response.PredictedAt,
	}, nil
}
