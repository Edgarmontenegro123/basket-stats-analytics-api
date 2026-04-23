package handlers

import "net/http"

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(`{"status":"ok","service":"basket-stats-analytics-api"}`))
	if err != nil {
		http.Error(w, "error writing response", http.StatusInternalServerError)
		return
	}
}
