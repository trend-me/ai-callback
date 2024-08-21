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
		var errParsed *exceptions.ErrorType
		if errors.As(err, &errParsed) && errParsed.Code == exceptions.PromptRoadMapNotFoundErrorCode {
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
		slog.DebugContext(ctx, "useCase.Handle",
			slog.String("details", "process finished - output published"))
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
		slog.String("details", "process finished - next step published"))
	return nil
}

func NewUseCase(
	queueAiPromptBuilder interfaces.QueueAiPromptBuilder,
	apiPromptRoadMapConfig interfaces.ApiPromptRoadMapConfig,
	apiPromptRoadMapConfigExecution interfaces.ApiPromptRoadMapConfigExecution,
	queueOutput interfaces.QueueOutput,
) interfaces.UseCase {
	return &UseCase{
		queueAiPromptBuilder:            queueAiPromptBuilder,
		apiPromptRoadMapConfigExecution: apiPromptRoadMapConfigExecution,
		apiPromptRoadMapConfig:          apiPromptRoadMapConfig,
		queueOutput:                     queueOutput,
	}
}
