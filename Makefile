MAKEFLAGS := --no-print-directory --silent

default: help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\._-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

t: test
test: fmt ## Run unit tests, alias: t
	cd core && go test ./... -timeout=60s -parallel=10 --cover

fmt: ## Format go code
	cd core && go mod tidy
	cd core && gofumpt -l -w .

tools: ## Install extra tools for development
	go install mvdan.cc/gofumpt@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: ## Lint the code locally
	cd core && golangci-lint run --timeout 600s

build: ## Builds the core of the game
	cd core && GOOS=js GOARCH=wasm go build -o ../ui/core.wasm ./cmd/core/main.go
