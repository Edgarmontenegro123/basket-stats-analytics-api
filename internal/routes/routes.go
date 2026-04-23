package routes

import (
	"net/http"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", handlers.HealthHandler)

	mux.HandleFunc("/uploads", handlers.UploadsHandler)
	mux.HandleFunc("/uploads/", handlers.UploadByIDHandler)

	mux.HandleFunc("/analytics/process", handlers.ProcessAnalyticsHandler)
	mux.HandleFunc("/analytics/games/", handlers.GamePlayerStatsHandler)
}
