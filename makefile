.PHONY: default run build test docs clean
# Variables
APP_NAME=kmk-cli
CLI_PATH=cmd/cli/main.go

# Tasks
default: run

run:
	@go run $(CLI_PATH)
build:
	@go build -o $(APP_NAME) $(CLI_PATH)
test:
	@go test ./ ...

clean:
	@rm -f $(APP_NAME)