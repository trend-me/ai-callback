package interfaces

import (
	"context"
)

type QueueAiCallbackConsumer interface {
	Consume(ctx context.Context) (chan error, error)
}
