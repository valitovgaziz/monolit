GOOSE_DRIVER=postgres
GOOSE_DBPASSWOR=postgres
GOOSE_DBSTRING="postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
GOOSE_MIGRATION_DIR=migrations

v:
	@goose -version

status:
	@goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} status

migration: run
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} up

.DEFAULT_GOAL=migration