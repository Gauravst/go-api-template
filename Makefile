include .env
export

# template migrate Variables
OLD_IMPORT=github.com/gauravst/got
OLD_CMD_DIR_API=cmd/got-api
OLD_CMD_DIR_CLI=cmd/got-cli

# info Variables
PROJECT_NAME=got
PROJECT_NAME_API=got-api
PROJECT_NAME_CLI=got-cli
GITHUB_USERNAME=gauravst

# variables
BINARY_NAME=got
GO_FILES=$(shell find . -name '*.go' -not -path './vendor/*')
MIGRATE_PATH=./migrations
DB_URL=$(DATABASE_URI)
APP_NAME=got
DOCKER_IMAGE_NAME=got
DOCKER_TAG=latest
PORT=8080

# default target
all: build

# build the application
build:
	@echo "Building the application..."
	go build -o bin/$(BINARY_NAME) cmd/$(PROJECT_NAME)/main.go

# run the api application
run-api:
	@echo "Running the application..."
	go run cmd/$(PROJECT_NAME_API)/main.go

# run the cli application
run-cli:
	@echo "Running the application..."
	go run cmd/$(PROJECT_NAME_CLI)/main.go

# run tests
test:
	@echo "Running tests..."
	go test -v ./...

# format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf bin/

# install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy

# run all up migrations
migrate-up:
	migrate -path $(MIGRATE_PATH) -database $(DB_URL) up

# run all down migrations
migrate-down:
	migrate -path $(MIGRATE_PATH) -database $(DB_URL) down

## build the Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) .

## run the Docker container
docker-run:
	docker run -p $(PORT):$(PORT) $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

## stop the Docker container
docker-stop:
	docker stop $$(docker ps -q --filter ancestor=$(DOCKER_IMAGE_NAME):$(DOCKER_TAG))

## remove the Docker container
docker-rm:
	docker rm $$(docker ps -a -q --filter ancestor=$(DOCKER_IMAGE_NAME):$(DOCKER_TAG))

## remove the Docker image
docker-rmi:
	docker rmi $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

## clean up Docker resources (stop, remove container, and remove image)
docker-clean: docker-stop docker-rm docker-rmi

## build and run the Go application
all: build run

## build and run the Docker container
docker-all: docker-build docker-run

# run all checks (format, test, build)
check: fmt test build

# template migrate 
SHELL := bash

# setup for project
setup:
	@bash -c '\
		echo "Enter your GitHub username :-"; \
		read USERNAME; \
		echo "Enter your project name :-"; \
		read PROJECT; \
		echo "Do you need API App? (y/n) :-"; \
		read API_APP; \
		echo "Do you need CLI App? (y/n) :-"; \
		read CLI_APP; \
		echo "Do you need Database? (y/n) :-"; \
		read DB; \
		if [ "$$DB" = "y" ]; then \
			echo "postgres or sqlite? (p/s) :-"; \
			read DB_TYPE; \
			echo "Do you need ORM (Gorm)? (y/n) :-"; \
			read ORM; \
		fi; \
		echo "Do you need unit/e2e/integration Testing? (y/n) :-"; \
		read TESTING; \
		echo "Do You need github workflow, CI/CD (y/n) :-"; \
		read WORKFLOW; \
		echo "Do you need Docker file setup? (y/n) :-"; \
		read IS_DOCKER; \
		LOWER_USERNAME=$$(echo $$USERNAME | tr "[:upper:]" "[:lower:]"); \
		LOWER_PROJECT=$$(echo $$PROJECT | tr "[:upper:]" "[:lower:]"); \
		NEW_IMPORT=github.com/$$LOWER_USERNAME/$$LOWER_PROJECT; \
		NEW_CMD_DIR=cmd/$$LOWER_PROJECT; \
		if [ "$$API_APP" = "n" ]; then \
			echo "Removing API App..."; \
			rm -rf $(OLD_CMD_DIR_API); \
			rm -rf internal/api; \
		else \
			echo "Renaming API App main.go..."; \
			mkdir -p $$NEW_CMD_DIR-api; \
			mv $(OLD_CMD_DIR_API)/main.go $$NEW_CMD_DIR-api/main.go; \
			rm -rf $(OLD_CMD_DIR_API); \
		fi; \
		if [ "$$CLI_APP" = "n" ]; then \
			echo "Removing CLI App..."; \
			rm -rf $(OLD_CMD_DIR_CLI); \
			rm -rf internal/cli; \
		else \
			echo "Renaming CLI App main.go..."; \
			mkdir -p $$NEW_CMD_DIR-cli; \
			mv $(OLD_CMD_DIR_CLI)/main.go $$NEW_CMD_DIR-cli/main.go; \
			rm -rf $(OLD_CMD_DIR_CLI); \
		fi; \
		if [ "$$DB" = "n" ]; then \
			echo "Removing database..."; \
			rm -rf internal/database; \
		else \
			if [ "$$ORM" = "y" ]; then \
				echo "Adding ORM..."; \
				rm -rf migrations; \
				rm -f internal/database/database.go; \
			else \
				echo "Adding Raw SQL setup"
				rm -f internal/database/init.go; \
				rm -f internal/database/postgres.go; \
				rm -f internal/database/sqlite.go; \
			fi; \
			if [ "$$DB_TYPE" = "p" ]; then \
				if [ "$$ORM" = "y" ]; then \
					echo "Adding Postgres Db..."; \
					rm -f internal/database/database.go; \
				else \
					# something here
				fi; \
			else \
				if [ "$$ORM" = "y" ]; then \
					echo "Adding Sqlite"
				else \
					# something here
				fi; \
			fi; \
		fi; \
		if [ "$$TESTING" = "n" ]; then \
			echo "Removing Testing..."; \
			rm -rf test; \
			find . -type f -name '*_test.go' -delete; \
		fi; \
		if [ "$$WORKFLOW" = "n" ]; then \
			echo "Removing workflow..."; \
			rm -rf .github/workflows; \
		else \
			if [ "$$TESTING" = "n" ]; then \
				rm -f .github/workflows/test.yml; \
			fi; \
		fi; \
		if [ "$$IS_DOCKER" = "n" ]; then \
			echo "Removing Docker setup..."; \
			rm -f Dockerfile; \
			rm -f .dockerignore; \
		fi; \
		echo "Replacing imports..."; \
		find . -type f -name "*.go" -exec sed -i "s|$(OLD_IMPORT)|$$NEW_IMPORT|g" {} +; \
		find . -type f -name "go.mod" -exec sed -i "s|$(OLD_IMPORT)|$$NEW_IMPORT|g" {} +; \
		find . -type f -name "go.sum" -exec sed -i "s|$(OLD_IMPORT)|$$NEW_IMPORT|g" {} +; \
		echo "Updating Makefile..."; \
		sed -i "s|OLD_IMPORT=$(OLD_IMPORT)|OLD_IMPORT=$$NEW_IMPORT|g" Makefile; \
		sed -i "s|OLD_CMD_DIR_API=$(OLD_CMD_DIR_API)|OLD_CMD_DIR_API=$${NEW_CMD_DIR}-api|g" Makefile; \
		sed -i "s|OLD_CMD_DIR_CLI=$(OLD_CMD_DIR_CLI)|OLD_CMD_DIR_CLI=$${NEW_CMD_DIR}-cli|g" Makefile; \
		sed -i "s|PROJECT_NAME=$(PROJECT_NAME)|PROJECT_NAME=$$LOWER_PROJECT|g" Makefile; \
		sed -i "s|PROJECT_NAME_API=$(PROJECT_NAME_API)|PROJECT_NAME_API=$${LOWER_PROJECT}-api|g" Makefile; \
		sed -i "s|PROJECT_NAME_CLI=$(PROJECT_NAME_CLI)|PROJECT_NAME_CLI=$${LOWER_PROJECT}-cli|g" Makefile; \
		sed -i "s|GITHUB_USERNAME=$(GITHUB_USERNAME)|GITHUB_USERNAME=$$LOWER_USERNAME|g" Makefile; \
		sed -i "s|BINARY_NAME=$(BINARY_NAME)|BINARY_NAME=$$LOWER_PROJECT|g" Makefile; \
		sed -i "s|APP_NAME=$(APP_NAME)|APP_NAME=$$LOWER_PROJECT|g" Makefile; \
		sed -i "s|DOCKER_IMAGE_NAME=$(DOCKER_IMAGE_NAME)|DOCKER_IMAGE_NAME=$$LOWER_PROJECT|g" Makefile; \
		echo "Setup completed!"; \
	'

# Help (list all targets)
help:
	@echo "Available targets:"
	@echo "  build    - Build the application"
	@echo "  run      - Run the application"
	@echo "  test     - Run tests"
	@echo "  fmt      - Format code"
	@echo "  clean    - Clean build artifacts"
	@echo "  deps     - Install dependencies"
	@echo "  check    - Run all checks (format, test, build)"
	@echo "  help     - Show this help message"
	@echo "  docker-build  - Build the Docker image"
	@echo "  docker-run    - Run the Docker container"
	@echo "  docker-stop   - Stop the Docker container"
	@echo "  docker-rm     - Remove the Docker container"
	@echo "  docker-rmi    - Remove the Docker image"
	@echo "  docker-clean  - Clean up Docker resources (stop, remove container, and remove image)"
	@echo "  docker-all    - Build and run the Docker container"
	@echo "  setup     - To Setup Project"
