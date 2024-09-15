package main

import (
	"os"
	"log/slog"

	"github.com/profclems/go-dotenv"
)


func main() {
	slog.Info("Init environment variable from .env file")
	dotenv.SetConfigFile("../.env")
	err := dotenv.Load()
	if err != nil {
		slog.Error("Can not read config file", "error", err)
		os.Exit(2)
	}

	slog.Info("Variables inited")
}