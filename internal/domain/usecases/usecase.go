package usecases

import (
	"context"
	"errors"
	"log/slog"

	"github.com/trend-me/ai-callback/internal/config/exceptions"
	"github.com/trend-me/ai-callback/internal/domain/interfaces"
	"github.com/trend-me/ai-callback/internal/domain/models"
)

type UseCase struct {
	apiPromptRoadMapConfig          interfaces.ApiPromptRoadMapConfig
	apiPromptRoadMapConfigExecution interfaces.ApiPromptRoadMapConfigExecution
	queueAiPromptBuilder            interfaces.QueueAiPromptBuilder
	queueOutput                     interfaces.QueueOutput
}

func (u UseCase) getNextPromptRoadMapStep(ctx context.Context, request *models.Request) (*models.PromptRoadMap, error) {
	nextPromptRoadMap, err := u.apiPromptRoadMapConfig.GetPromptRoadMap(ctx, request.PromptRoadMapConfigName, request.PromptRoadMapStep+1)
	if err != nil {
		var errParsed exceptions.ErrorType
		if errors.As(err, &errParsed) && errParsed.Code == exceptions.GetPromptRoadMapConfigErrorCode {
			return nil, nil
		}
		return nil, err
	}
	return nextPromptRoadMap, nil
}

func (u UseCase) Handle(ctx context.Context, request *models.Request) error {
	slog.InfoContext(ctx, "useCase.Handle",
		slog.String("details", "process started"))

	nextPromptRoadMap, err := u.getNextPromptRoadMapStep(ctx, request)
	if err != nil {
		return err
	}

	if nextPromptRoadMap == nil {
		err = u.queueOutput.Publish(ctx, request.OutputQueue, request)
		return err
	}

	if err = u.apiPromptRoadMapConfigExecution.UpdateStepInExecutionById(
		ctx, request.PromptRoadMapConfigExecutionId, nextPromptRoadMap.Step); err != nil {
		return err
	}

	request.PromptRoadMapStep++

	err = u.queueAiPromptBuilder.Publish(ctx, request)
	if err != nil {
		return err
	}

	slog.DebugContext(ctx, "useCase.Handle",
		slog.String("details", "process finished"))
	return nil
}

func NewUseCase(
	queueAiPromptBuilder interfaces.QueueAiPromptBuilder,
	apiPromptRoadMapConfig interfaces.ApiPromptRoadMapConfig,
	queueOutput interfaces.QueueOutput,
) interfaces.UseCase {
	return &UseCase{
		queueAiPromptBuilder:   queueAiPromptBuilder,
		apiPromptRoadMapConfig: apiPromptRoadMapConfig,
		queueOutput:            queueOutput,
	}
}
