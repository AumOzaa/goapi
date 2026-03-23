package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AumOzaa/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// adding the global middleware
	r.Use(chimiddle.StripSlashes) // this  handles the tariling '/' in the route

	r.Route("/account", func(router chi.Router) {
		fmt.Printf("During compilation :IN the account route\n")
		// can use this to define the get method

		router.Use(middleware.Authorization) // This adds the middleware , the .Use

		fmt.Println("During compilation : After the authorization")

		router.Get("/coins", GetCoinBalance) // The other argument takes the function with args

	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) { //adding more

		w.Header().Set("Content-Type", "application/json")
		var userResponse string = "done"
		json.NewEncoder(w).Encode(userResponse)
	})
}
