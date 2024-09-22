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
	close, err := logger.InitLogger()
	if err != nil {
		panic("can't init logger error = " + err.Error())
	}
	close()
	server.Middleware()
	server.Routing()
	server.Start()
	storage.DBconnection()
	slog.Info("Server is started seccesfully", "SERVER_PORT", dotenv.GetString("SERVER_PORT"))
	slog.Info("server is closed", "info", <-server.Done)
}
