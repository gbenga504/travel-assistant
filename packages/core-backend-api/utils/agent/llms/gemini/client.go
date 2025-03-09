package gemini

import (
	"context"

	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	client *genai.Client
}

func NewClient(ctx context.Context, apiKey string) *GeminiClient {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))

	// It is super important to have a client, hence if there is an error
	// then we use log fatal which returns an os exit 1 also
	if err != nil {
		logger.Fatal(err.Error(), logger.ErrorOpt{
			Name:    errors.Name(errors.ErrAIClientNotLoaded),
			Message: errors.Message(errors.ErrAIClientNotLoaded),
		})
	}

	return &GeminiClient{
		client,
	}
}

func (gc *GeminiClient) Close() {
	gc.client.Close()
}
