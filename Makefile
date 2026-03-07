SHELL := /bin/zsh
.DEFAULT_GOAL := run

# Ensures that make always runs the commands for these targets without checking for a file with the same name
.PHONY: fmt vet build run clean tidy

fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build -o bin/client .
run: build
	./bin/client
clean:
	go clean
	rm -rf bin/
tidy:
	go mod tidy
test:
	go test ./...
dev:
	go run main.go