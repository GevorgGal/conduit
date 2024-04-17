.PHONY: build test clean docker

# Binary output directory
BIN_DIR := ./bin
# Binary name
BINARY_NAME := conduit-connector-influxdb

# Default target
all: build

# Build the project
build:
	@echo "Building the project..."
	@go build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/conduit-connector-influxdb

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Clean up the project
clean:
	@echo "Cleaning up..."
	@go clean
	@rm -rf $(BIN_DIR)

# Build Docker container
docker:
	@echo "Building Docker image..."
	@docker build -t $(BINARY_NAME):latest .

# Start Docker Compose
docker-up: docker
	@echo "Starting Docker Compose services..."
	@docker-compose up -d

# Stop Docker Compose
docker-down:
	@echo "Stopping Docker Compose services..."
	@docker-compose down