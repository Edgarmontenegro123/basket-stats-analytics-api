package routes

import (
	"net/http"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", handlers.HealthHandler)
}
