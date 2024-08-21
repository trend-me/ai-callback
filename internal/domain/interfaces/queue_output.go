package interfaces

import (
	"context"

	"github.com/trend-me/ai-callback/internal/domain/models"
)

type QueueOutput interface {
	Publish(ctx context.Context, name string, request *models.Request) error
}
