.PHONY: default start run build dockerbuild test clean stop

# Variables
APP_NAME=goingcrazy
OS=linux
OS_ARCH=amd64

# Tasks
default: start

start: 
	@docker-compose up
run:
	@go run main.go
build:
	@rm -f $(APP_NAME)
	@env GOS=$(OS) GOARCH=$(OS_ARCH) go build -o $(APP_NAME) main.go
dockerbuild:
	@rm -f $(APP_NAME)
	@env GOS=$(OS) GOARCH=$(OS_ARCH) go build -o $(APP_NAME) main.go
	@docker-compose down -v
	@docker-compose build
	@docker-compose up
test:
	@go test ./ ...
clean:
	@rm -f $(APP_NAME)
stop:
	@docker-compose down