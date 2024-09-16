build:
	@go build -o bin/mon cmd/main.go

run: build
	@./bin/mon

test:
	@go test ./... -v

.DEFAULT_GOAL=run