GOOSE_DRIVER=postgres
GOOSE_DBPASSWOR=postgres
GOOSE_DBSTRING="user=postgres dbname=postgres sslmode=disable"
GOOSE_MIGRATION_DIR=migrations

build:
	@go build -o SetEnvVar/set SetEnvVar/main.go

run: build
	@./SetEnvVar/set

v:
	@goose -version

status:
	@goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} status

migration: run
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} up

.DEFAULT_GOAL=migration