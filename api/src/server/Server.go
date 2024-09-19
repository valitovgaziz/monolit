package server

import (
	"log/slog"
	"net/http"

	"github.com/profclems/go-dotenv"
)

// up server ${SERVER_PORT} (.env file) port on gorutin
func Start() {
	go func() {
		defer close(Done)
		serverPort := ":" + dotenv.GetString("SERVER_PORT")
		err := http.ListenAndServe(serverPort, r)
		if err != nil {
			slog.Error("Can't start server: ", "error", err)
		}
	}()
}
