package start

import (
	"log/slog"
	"net/http"
)

// up server on 3000 port on gorutin
func Server() {
	go func() {
		defer close(Done)
		err := http.ListenAndServe(":3000", r)
		if err != nil {
			slog.Error("Can't start server: ", "error", err)
		}

	}()
}
