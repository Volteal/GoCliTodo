# Go Commands
tidy:
	@go mod tidy

test:
	@go test -coverage ./tests/...

run:
	@go run ./cmd/todo-app/main.go

build:
	@go build -v -o ./bin/todo-app/todo ./cmd/todo-app/main.go

install:
	@go install ./cmd/todo