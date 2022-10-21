APP=stakefish
APP_VERSION:=$(shell git rev-parse HEAD)
APP_EXECUTABLE="./out/$(APP)"

build:  $(shell find . -type f -name '*.go')
	mkdir -p out/
	go build -ldflags "-X github.com/prybintsev/stakefish/internal/version.Version=$(APP_VERSION)" -o $(APP_EXECUTABLE) ./cmd

swagger:
	swag init -g cmd/main.go --parseDependency

test:
	go test -v ./tests/...
