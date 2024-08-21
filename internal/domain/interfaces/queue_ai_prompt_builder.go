package interfaces

import (
	"context"

	"github.com/trend-me/ai-callback/internal/domain/models"
)

type QueueAiPromptBuilder interface {
	Publish(ctx context.Context, request *models.Request) error
}
