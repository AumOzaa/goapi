package handlers

import (
	"github.com/AumOzaa/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// adding the global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		// can use this to define the get method

		router.Use(middleware.Authorization) // TODO: Need to add this later

		router.Get("/coins", GetCoinsBalance) // TODO: Define th GetCoinBalance function

	})
}
