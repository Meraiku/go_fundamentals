package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) SetupHandlers() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {

		r.Get("/", handleRoot)

		r.Get("/search", api.handleSearch)

	})

	return r
}
