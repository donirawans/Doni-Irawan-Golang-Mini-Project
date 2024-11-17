package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ResponseAI(ctx context.Context, question string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("AI_API_KEY")))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	modelAI := client.GenerativeModel("gemini-pro")
	modelAI.SetTemperature(0)

	resp, err := modelAI.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	answer := resp.Candidates[0].Content.Parts[0]
	answerString := fmt.Sprintf("%v", answer)

	// Bersihkan simbol tambahan dari teks
	answerString = strings.ReplaceAll(answerString, "*", "")
	answerString = strings.ReplaceAll(answerString, "**", "")
	answerString = strings.ReplaceAll(answerString, "\n\n", " -")
	return answerString, nil
}
