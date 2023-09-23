.PHONY: default run build test clean stop

# Variables
APP_NAME=goingcrazy

# Tasks
default: run

run:
	@docker-compose up -d
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
clean:
	@rm -f $(APP_NAME)
stop:
	@docker-compose down