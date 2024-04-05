APP_NAME=circle
VERSION?=0.0.1

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0
GO111MODULE=auto
GOVERSION=1.21.1

API_IMG_NAME=api
BIN_DIR=./bin

ARGS=$(filter-out $@,$(MAKECMDGOALS))

%:
	@:

.YELLOW := $(shell tput -Txterm setaf 3)
.RESET  := $(shell tput -Txterm sgr0)

.DEFAULT_GOAL := help

.PHONY: help

help:
	@echo "${APP_NAME} - v${VERSION}\n"
	@echo "Makefile targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo

# * GOLANG * #
## Clean built binaries directory
clean: 
	@echo "${.YELLOW}--- Go: clean ---${.RESET}"
	rm -rf $(BIN_DIR)

## Get dependencies
tidy: 
	@echo "${.YELLOW}--- Go: tidy ---${.RESET}"
	go mod tidy -compat=$(GOVERSION)

## Build binary
build: tidy
	@echo "${.YELLOW}--- Go: build ---${.RESET}"
	mkdir -p $(BIN_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) GO111MODULE=$(GO111MODULE) CGO_ENABLED=$(CGO_ENABLED) go build -o $(BIN_DIR) ./cmd/userapi

## Build local development binary
build-dev: 
	@echo "${.YELLOW}--- Go: build dev ---${.RESET}"
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR) ./cmd/userapi

## Run Go binary
run: docker-up
	@echo "${.YELLOW}--- Go: run ---${.RESET}"
	$(BIN_DIR)/userapi

## Run Go unit tests
tests: 
	@echo "${.YELLOW}--- Go: tests ---${.RESET}"
	go test -v -race ./...

## Run Go integration tests
tests-integration: 
	@echo "${.YELLOW}--- Go: tests ---${.RESET}"
	go test -tags e2e -v -race ./...

## Run Go integration tests
tests-script: 
	@echo "${.YELLOW}--- Manual tests script ---${.RESET}"
	./scripts/tests.sh

## Lint Go code
lint: 
	@echo "${.YELLOW}--- Go: lint ---${.RESET}"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run


# * DOCKER * #
## Build all docker images
docker-build: build
	@echo "${.YELLOW}--- Docker: build ---${.RESET}"
	docker build --build-arg SERVICE_NAME=$(USER_API_IMG_NAME) -t $(USER_API_IMG_NAME):latest .

## Build one docker image
docker-build-img: 
	@echo "${.YELLOW}--- Docker: build img ---${.RESET}"
	docker build --build-arg SERVICE_NAME=$(ARGS) -t $(ARGS):latest .

## Up docker-compose cluster
docker-up:
	@echo "${.YELLOW}--- Docker: up ---${.RESET}"
	docker-compose up -d --build --remove-orphans

## Down docker-compose cluster
docker-down: 
	@echo "${.YELLOW}--- Docker: down ---${.RESET}"
	docker-compose down --remove-orphans

# * KAFKA * #
kafka-create-topic:
	@echo "${.YELLOW}--- Kafka: create topic ---${.RESET}"
	docker-compose exec kafka kafka-topics.sh --create --topic $(ARGS) --partitions 1 --replication-factor 1 --if-not-exists --bootstrap-server kafka:9092
