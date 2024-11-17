package repository

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
	"warningfloodsystem/domain/model"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIWeatherRepository interface {
	GetAIResponse(ctx context.Context, question string, curahHujan, tinggiSungai float64) (*model.AIWeather, error)
}

type aiWeatherRepository struct{}

func NewAIWeatherRepository() AIWeatherRepository {
	return &aiWeatherRepository{}
}

func (r *aiWeatherRepository) GetAIResponse(ctx context.Context, question string, curahHujan, tinggiSungai float64) (*model.AIWeather, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("AI_API_KEY")))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize AI client: %w", err)
	}

	modelAI := client.GenerativeModel("gemini-pro")
	modelAI.SetTemperature(0.7)

	resp, err := modelAI.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		return nil, fmt.Errorf("error generating AI response: %w", err)
	}

	if len(resp.Candidates) == 0 {
		return nil, fmt.Errorf("AI response contains no candidates")
	}

	// Gabungkan semua bagian dari respons
	fullResponse := ""
	for _, candidate := range resp.Candidates {
		if candidate.Content != nil {
			for _, part := range candidate.Content.Parts {
				fullResponse += fmt.Sprintf("%v", part) + " "
			}
		}
	}
	fullResponse = strings.TrimSpace(fullResponse)

	return &model.AIWeather{
		Rekomendasi: fullResponse,
		PredictedAt: time.Now(),
	}, nil
}
