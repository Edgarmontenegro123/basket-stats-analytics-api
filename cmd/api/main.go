package main

import (
	"log"
	"net/http"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/routes"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	server := &http.Server{
		Addr:    ":8081",
		Handler: enableCORS(mux),
	}

	log.Println("Analytics API running on http://localhost:8081")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
