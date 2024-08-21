//go:build wireinject

package injectors

import (
	"github.com/google/wire"
	"github.com/trend-me/ai-callback/internal/config/properties"
	"github.com/trend-me/ai-callback/internal/delivery/controllers"
	"github.com/trend-me/ai-callback/internal/domain/interfaces"
	"github.com/trend-me/ai-callback/internal/domain/usecases"
	"github.com/trend-me/ai-callback/internal/integration/api"
	"github.com/trend-me/ai-callback/internal/integration/connections"
	"github.com/trend-me/ai-callback/internal/integration/queues"
	"github.com/trend-me/golang-rabbitmq-lib/rabbitmq"
)

func newQueueConnectionOutputGetter(connection *rabbitmq.Connection) queues.ConnectionOutputGetter {
	return func(queueName string) queues.ConnectionOutput {
		return rabbitmq.NewQueue(
			connection,
			queueName,
			rabbitmq.ContentTypeJson,
			properties.CreateQueueIfNX(),
			true,
			true,
		)
	}
}

func newQueueConnectionAiCallbackConsumer(connection *rabbitmq.Connection) queues.ConnectionAiCallbackConsumer {
	return rabbitmq.NewQueue(
		connection,
		properties.QueueAiRequester,
		rabbitmq.ContentTypeJson,
		properties.CreateQueueIfNX(),
		true,
		true,
	)
}

func newQueueConnectionAiRequester(connection *rabbitmq.Connection) queues.ConnectionAiPromptBuilder {
	return rabbitmq.NewQueue(
		connection,
		properties.QueueAiCallback,
		rabbitmq.ContentTypeJson,
		properties.CreateQueueIfNX(),
		true,
		true,
	)
}

func urlApiPromptRoadMapConfigGetter() api.UrlApiPromptRoadMapConfig {
	return properties.UrlApiPromptRoadMapConfig
}

func urlApiPromptRoadMapConfigExecutionGetter() api.UrlApiPromptRoadMapConfigExecution {
	return properties.UrlApiPromptRoadMapConfigExecution
}

func InitializeQueueAiRequesterConsumer() (interfaces.QueueAiCallbackConsumer, error) {
	wire.Build(
		controllers.NewController,
		urlApiPromptRoadMapConfigGetter,
		api.NewPromptRoadMapConfig,
		urlApiPromptRoadMapConfigExecutionGetter,
		api.NewPromptRoadMapConfigExecution,
		usecases.NewUseCase,
		newQueueConnectionOutputGetter,
		queues.NewOutput,
		queues.NewAiPromptBuilder,
		connections.ConnectQueue,
		newQueueConnectionAiCallbackConsumer,
		newQueueConnectionAiRequester,
		queues.NewAiCallbackConsumer)
	return nil, nil
}
