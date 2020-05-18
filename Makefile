.DEFAULT_GOAL := build

.PHONY: build clean go_test_unit goose_up postgres_up
.PHONY: docker_build_image
.PHONY: application_start application_test run_test_integration

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


# Docker targets.
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

application_start: postgres_up
	@echo ">>> Sleeping 10 seconds until postgres start."
	@sleep 10
	@echo ">>> Starting application."
	@docker-compose $(DOCKER_COMPOSE_OPTIONS) up -d $(SERVICE)

goose_up:
	@docker-compose $(DOCKER_COMPOSE_OPTIONS) run \
            --rm \
            -v $$PWD:/app \
			-w /app \
            goose-migrate

# Go targets.
go_get:
	@echo '>>> Getting go modules.'
	@env GOPROXY=$(GOPROXY) go mod download
	
go_test_unit:
	@echo ">>> Running unit tests."
	@env GOPROXY=$(GOPROXY) go test -v -tags unit -cover ./... -coverprofile=coverunit.out

go_test_integration:
	@echo ">>> Running integrartion tests."
	@env GOPROXY=$(GOPROXY) go test -v -tags="integration" ./test/integration/...


run_test_integration:
	@echo ">>> Running tests"
	@docker-compose $(DOCKER_COMPOSE_OPTIONS) run \
    		--rm \
    		-v $$PWD:/go/src/$(GO_IMPORT_PATH) \
    		-w /go/src/$(GO_IMPORT_PATH) \
    		--no-deps \
    		-e GO111MODULE=$(GO111MODULE) \
    		integration-tests

application_test: docker_down application_start goose_up  run_test_integration 	docker_down		
