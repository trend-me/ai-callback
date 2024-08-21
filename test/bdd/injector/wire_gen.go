// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/trend-me/ai-callback/internal/config/connections"
	"github.com/trend-me/ai-callback/internal/config/properties"
	"github.com/trend-me/ai-callback/internal/delivery/controllers"
	"github.com/trend-me/ai-callback/internal/domain/factories"
	"github.com/trend-me/ai-callback/internal/domain/interfaces"
	"github.com/trend-me/ai-callback/internal/domain/usecases"
	"github.com/trend-me/ai-callback/internal/integration/api"
	"github.com/trend-me/ai-callback/internal/integration/queues"
	"github.com/trend-me/golang-rabbitmq-lib/rabbitmq"
)

// Injectors from wire.go:

func InitializeQueueAiRequesterConsumerMock(geminiMock interfaces.Ai) (interfaces.QueueAiCallbackConsumer, error) {
	connection, err := connections.ConnectQueue()
	if err != nil {
		return nil, err
	}
	connectionAiCallback := newQueueConnectionAiCallback(connection)
	queueAiCallback := queues.NewAiPromptBuilder(connectionAiCallback)
	aiFactory := factories.NewAiFactory(geminiMock)
	urlApiPromptRoadMapConfig := urlApiPromptRoadMapConfigGetter()
	apiPromptRoadMapConfig := api.NewApiPromptRoadMapConfig(urlApiPromptRoadMapConfig)
	urlApiValidation := urlApiValidationGetter()
	apiValidation := api.NewValidation(urlApiValidation)
	useCase := usecases.NewUseCase(queueAiCallback, aiFactory, apiPromptRoadMapConfig, apiValidation)
	controller := controllers.NewController(useCase)
	connectionAiRequesterConsumer := newQueueConnectionAiRequesterConsumer(connection)
	queueAiRequesterConsumer := newQueueAiRequesterConsumer(controller, connectionAiRequesterConsumer)
	return queueAiRequesterConsumer, nil
}

// wire.go:

func newQueueConnectionAiRequesterConsumer(connection *rabbitmq.Connection) queues.ConnectionAiRequesterConsumer {
	return rabbitmq.NewQueue(
		connection, properties.QueueAiRequester, rabbitmq.ContentTypeJson, properties.CreateQueueIfNX(), true,
		true,
	)
}

func newQueueConnectionAiCallback(connection *rabbitmq.Connection) queues.ConnectionAiCallback {
	return rabbitmq.NewQueue(
		connection, properties.QueueAiCallback, rabbitmq.ContentTypeJson, properties.CreateQueueIfNX(), true,
		true,
	)
}

func newQueueAiRequesterConsumer(controller interfaces.Controller, connectionAiPromptBuilderConsumer queues.ConnectionAiRequesterConsumer) interfaces.QueueAiCallbackConsumer {
	return queues.NewAiPromptBuilderConsumer(connectionAiPromptBuilderConsumer, controller)
}

func urlApiValidationGetter() api.UrlApiValidation {
	return properties.UrlApiValidation
}

func urlApiPromptRoadMapConfigGetter() api.UrlApiPromptRoadMapConfig {
	return properties.UrlApiPromptRoadMapConfig
}
