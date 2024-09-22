package main

import (
	"api/src/logger"
	"api/src/server"
	"api/src/storage"
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
	logger.InitLogger()
	server.Middleware()
	server.Routing()
	server.Start()
	storage.DBconnection()
	slog.Info("Server is started seccesfully", "SERVER_PORT", dotenv.GetString("SERVER_PORT"))
	slog.Info("server is closed", "info", <-server.Done)
}
