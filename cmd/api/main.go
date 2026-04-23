package main

import (
	"log"
	"net/http"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	log.Println("Analytics API running on http://localhost:8081")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
