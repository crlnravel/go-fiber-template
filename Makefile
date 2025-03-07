.PHONY: clean critic security lint test build run

APP_NAME = minecart
MAIN_DIR = $(PWD)/cmd/app
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_DIR)

run:
	go run $(MAIN_DIR)

gen.docs:
	swag init -d "./cmd/app" --pd --pdl 3

