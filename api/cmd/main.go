package main

import (
	"api/src/start"
	"log/slog"
	"os"

	"github.com/profclems/go-dotenv"
)

func init() {
	dotenv.SetConfigFile(".env")
	err := dotenv.Load()
	if err != nil {
		slog.Error("Can not read config file", "error", err)
		os.Exit(2)
	}
}

func main() {
	start.Middleware()
	start.Routing()
	start.Server()
	start.DBconnection()
	slog.Info("server is closed", "info", <-start.Done)
}
