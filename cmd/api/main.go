package main

import (
	"fmt"
	"net/http"

	"github.com/avukadin/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus" // will be aliased as log
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter() // struct that'd be used to make the API
	handlers.Handler(r)

	fmt.Println("Starting GO API service....")

	err := http.ListenAndServe("localhost:8000", r) //starting the server

	if err != nil {
		log.Error(err)
	}
}
