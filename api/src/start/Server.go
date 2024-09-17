package start

import (
	"api/src/handlers/auth"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var Done = make(chan bool)
var r = chi.NewRouter()

func InitRouting() {
	slog.Info("Init routing")

	// middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.RequestID)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.NoCache)
	r.Use(middleware.Recoverer)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	// public Routes
	r.Group(func(r chi.Router) {
		r.Post("/signup", auth.Register) // register
		r.Post("/signin", auth.Login)    // signin
		// r.Get("/search", srch.Search)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hi"))
			w.WriteHeader(http.StatusOK)
		})
	})

}

// up server on 3000 port on gorutin
func StartServer() {
	go func() {
		defer close(Done)
		err := http.ListenAndServe(":3000", r)
		if err != nil {
			slog.Error("Can't start server: ", "error", err)
		}

	}()
}
