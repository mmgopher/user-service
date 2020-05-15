.DEFAULT_GOAL := build

.PHONY: build clean go_test_unit goose_up postgres_up

## Project variables

SERVICE=user-service
GOPROXY=https://proxy.golang.org
PATH_DOCKER_FILE=$(realpath ./build/Dockerfile)
DOCKER_COMPOSE_FILE_PATH=$(realpath ./build/docker-compose.yml)
DOCKER_COMPOSE_OPTIONS= -f $(DOCKER_COMPOSE_FILE_PATH)
GO111MODULE=on
GO_IMPORT_PATH=$(shell go list .)

build:
	env GOOS=linux GOPROXY=$(GOPROXY) go build -ldflags="-s -w" -o bin/$(SERVICE)

clean:
	rm -rf ./bin

docker_build_image:
	@echo ">>> Building docker image with service binary."
	docker build \
		-t $(SERVICE) \
		--build-arg GOPROXY=$(GOPROXY) \
		--build-arg GO111MODULE=$(GO111MODULE) \
		--build-arg GO_IMPORT_PATH=$(GO_IMPORT_PATH) \
		-f $(PATH_DOCKER_FILE) \
		.


docker_down:
	@docker-compose  $(DOCKER_COMPOSE_OPTIONS) down -v --remove-orphans

postgres_up:
	@docker-compose $(DOCKER_COMPOSE_OPTIONS) up \
	-d \
	user-service-postgres 

goose_up:
	@docker-compose $(DOCKER_COMPOSE_OPTIONS) run \
            --rm \
            -v $$PWD:/app \
			-w /app \
            goose-migrate


go_get:
	@echo '>>> Getting go modules.'
	@env GOPROXY=$(GOPROXY) GOPRIVATE=$(GOPRIVATE) go mod download
	
go_test_unit:
	@echo ">>> Running unit tests."
	@env GIN_MODE=release GOPROXY=$(GOPROXY) go test -v -tags unit -cover ./... -coverprofile=coverunit.out
