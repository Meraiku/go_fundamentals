package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupHandlers() *chi.Mux {

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello world"))
		})
	})

	return r

}
