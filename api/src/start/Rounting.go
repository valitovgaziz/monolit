package start

import (
	"api/src/handlers/auth"
	"net/http"

	"github.com/go-chi/chi"
)

func Routing() {

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
