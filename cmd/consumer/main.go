package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/trend-me/ai-callback/internal/config/injectors"
	"github.com/trend-me/ai-callback/internal/integration/connections"
)

func main() {
	ctx := context.Background()

	// Initialize consumer
	consumer, err := injectors.InitializeQueueAiCallbackConsumer()
	if err != nil {
		log.Fatalf("Error initializing consumer: %v", err)
		return
	}

	errChan, err := consumer.Consume(ctx)
	if err != nil {
		log.Fatalf("Error initializing consumer: %v", err)
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		if err != nil {
			log.Printf("Error consuming messages: %v", err)
		}
	case sig := <-sigChan:
		log.Printf("Received signal: %v. Shutting down...", sig)
	}

	defer connections.Disconnect()
}
