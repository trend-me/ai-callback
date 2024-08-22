package queues

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/trend-me/ai-callback/internal/config/exceptions"
	"github.com/trend-me/ai-callback/internal/domain/interfaces"
	"github.com/trend-me/ai-callback/internal/domain/models"
)

type (
	ConnectionAiPromptBuilder interface {
		Publish(ctx context.Context, b []byte) (err error)
		Connect() (err error)
	}

	aiRequesterMessage struct {
		PromptRoadMapConfigName        string         `json:"prompt_road_map_config_name"`
		PromptRoadMapStep              int            `json:"prompt_road_map_step"`
		PromptRoadMapConfigExecutionId string         `json:"prompt_road_map_config_execution_id"`
		OutputQueue                    string         `json:"output_queue"`
		Model                          string         `json:"model"`
		Metadata                       map[string]any `json:"metadata"`
	}

	aiPromptBuilder struct {
		queue ConnectionAiPromptBuilder
	}
)

func (a aiPromptBuilder) Publish(ctx context.Context, request *models.Request) error {
	slog.InfoContext(ctx, "aiPromptBuilder.Publish",
		slog.String("details", "process started"))

	b, err := json.Marshal(aiRequesterMessage{
		PromptRoadMapConfigName:        request.PromptRoadMapConfigName,
		PromptRoadMapConfigExecutionId: request.PromptRoadMapConfigExecutionId,
		PromptRoadMapStep:              request.PromptRoadMapStep,
		OutputQueue:                    request.OutputQueue,
		Model:                          request.Model,
		Metadata:                       request.Metadata,
	})
	if err != nil {
		return exceptions.NewValidationError("error parsing ai-requester message", err.Error())
	}

	err = a.queue.Publish(ctx, b)
	if err != nil {
		return exceptions.NewQueueError(err.Error())
	}

	return nil
}

func NewAiPromptBuilder(queue ConnectionAiPromptBuilder) interfaces.QueueAiPromptBuilder {
	return &aiPromptBuilder{
		queue: queue,
	}
}
