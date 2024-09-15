build:
	@go build -o bin/InitEnvVariable cmd/Main.go

run: build
	@./bin/InitEnvVariable

v:
	@goose -version

status:
	@goose -dir migration ${GOOSE_DRIVER} ${GOOSE_DBSTRING} status

migration:
	@goose -dir migration ${GOOSE_DRIVER} ${GOOSE_DBSTRING} up

.DEFAULT_GOAL=migration