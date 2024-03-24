# Simple Makefile for a Go project

# Build the application
build:
	@echo "Building..."

	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."

	@go test -v $(shell go list ./... | grep -v /test/)

test-cover:
	@go test -coverprofile cover.out -v $(shell go list ./... | grep -v /test/)
	
	@go tool cover -html=cover.out

gen-mocks:
	@mockery --all --with-expecter --keeptree

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

