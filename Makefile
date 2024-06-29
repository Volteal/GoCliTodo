# Go Commands
tidy:
	@go mod tidy

test:
	@go test -coverage ./tests/...

run:
	@go run ./cmd/todo/main.go

build:
	@go build -v -o ./bin/todo ./cmd/todo/main.go

install:
	@go install ./cmd/todo