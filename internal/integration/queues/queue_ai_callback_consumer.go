package queues

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/trend-me/ai-callback/internal/domain/interfaces"
)

type (
	ConnectionAiCallbackConsumer interface {
		Consume(ctx context.Context, handler func(delivery amqp.Delivery) error) (chan error, error)
	}

	aiRequesterConsumer struct {
		queue      ConnectionAiCallbackConsumer
		controller interfaces.Controller
	}
)

func (a aiRequesterConsumer) Consume(ctx context.Context) (chan error, error) {
	return a.queue.Consume(ctx, a.controller.Handle)
}

func NewAiCallbackConsumer(queue ConnectionAiCallbackConsumer, controller interfaces.Controller) interfaces.QueueAiCallbackConsumer {
	return &aiRequesterConsumer{queue: queue, controller: controller}
}
