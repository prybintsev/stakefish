APP=stakefish
APP_VERSION:=$(shell git describe --tags)
APP_EXECUTABLE="./out/$(APP)"

build:  $(shell find . -type f -name '*.go')
	mkdir -p out/
	go build -ldflags "-X github.com/prybintsev/stakefish/internal/version.Version=$(APP_VERSION)" -o $(APP_EXECUTABLE) ./cmd
