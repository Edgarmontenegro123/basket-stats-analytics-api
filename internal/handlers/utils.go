package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "error generating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "error writing response", http.StatusInternalServerError)
		return
	}
}

func generateID() string {
	return time.Now().Format("20060102150405.000000")
}
