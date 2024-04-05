APP_NAME=circle
VERSION?=0.0.1

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0
GO111MODULE=auto
GOVERSION=1.21.1

BIN_DIR=./bin

API_APP_PATH=./go/cmd/api
KAFKA_APP_PATH=./go/cmd/kafka

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
	GOOS=$(GOOS) GOARCH=$(GOARCH) GO111MODULE=$(GO111MODULE) CGO_ENABLED=$(CGO_ENABLED) go build -o $(BIN_DIR) .go/cmd/**

## Run Go API
run-api:
	@echo "${.YELLOW}--- Go: run api app ---${.RESET}"
	go run $(API_APP_PATH) -sql ./db/create_table_TB01.sql


## Run Go kafka app
run-kafka:
	@echo "${.YELLOW}--- Go: run kafka app ---${.RESET}"
	go run $(KAFKA_APP_PATH) -m message -k key

## Run Go unit tests
tests: 
	@echo "${.YELLOW}--- Go: tests ---${.RESET}"
	go test -v -race ./...

## Run integration test
test-script: 
	@echo "${.YELLOW}--- Manual test script ---${.RESET}"
	./scripts/test.sh

# * DOCKER * #
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
	docker-compose exec kafka kafka-topics.sh --create --topic $(ARGS) --partitions 3 --replication-factor 1 --if-not-exists --bootstrap-server kafka:9092
