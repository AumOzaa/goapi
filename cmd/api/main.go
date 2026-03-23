package main

import (
	"fmt"
	"net/http"

	"github.com/AumOzaa/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus" // will be aliased as log
)

func main() {
	log.SetReportCaller(true)
	fmt.Println("COMPILATION : In the main.go 1")
	var r *chi.Mux = chi.NewRouter() // struct that'd be used to make the API
	handlers.Handler(r)

	fmt.Println("COMPILATION : Printing after the route")
	fmt.Println("Starting GO API service....")

	err := http.ListenAndServe("localhost:8000", r) //starting the server

	fmt.Println("In the main.go file")

	if err != nil {
		log.Error(err)
	}
}
