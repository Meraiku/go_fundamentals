package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupHandlers() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {

		r.Get("/", handleRoot)

	})

	return r
}
