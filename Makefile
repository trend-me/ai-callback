build:
	go mod tidy
	go build -o consumer.out ./cmd/consumer

wire:
	go install github.com/google/wire/cmd/wire
	wire  ./internal/config/injectors

wire_mock:
	go install github.com/google/wire/cmd/wire
	wire  ./test/bdd/injectors

bdd: 
	go test -v ./test/bdd/steps/step_definitions_test.go