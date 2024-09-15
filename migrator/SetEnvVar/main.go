package main

import (
	"log/slog"
	"os"

	"github.com/profclems/go-dotenv"
)

func main() {
	slog.Info("Set environment variable from .env file.")
	dotenv.SetConfigFile(".env")
	err := dotenv.Load()
	if err != nil {
		slog.Error("Can not read config file", "error", err)
		os.Exit(2)
	}
	os.Setenv("GOOSE_DRIVER", dotenv.GetString("GOOSE_DRIVER"))
	os.Setenv("GOOSE_DBPASSWORD", dotenv.GetString("GOOSE_DBPASSWORD"))
	os.Setenv("GOOSE_DBSTRING", dotenv.GetString("GOOSE_DBSTRING"))
	os.Setenv("GOOSE_MIGRATION_DIR", dotenv.GetString("GOOSE_MIGRATION_DIR"))
	slog.Info("Environment variable for goose is set.")
}
