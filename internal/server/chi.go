package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupHandlers() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello world"))
		})
	})

	return r

}
