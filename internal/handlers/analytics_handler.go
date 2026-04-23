package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/models"
)

var playerStats []models.PlayerStats

type processAnalyticsRequest struct {
	UploadID string `json:"upload_id"`
}

func ProcessAnalyticsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		processAnalytics(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func GamePlayerStatsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listPlayerStatsByGameID(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func processAnalytics(w http.ResponseWriter, r *http.Request) {
	var req processAnalyticsRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.UploadID == "" {
		http.Error(w, "upload_id is required", http.StatusBadRequest)
		return
	}

	upload, found := findUploadByID(req.UploadID)
	if !found {
		http.Error(w, "upload not found", http.StatusNotFound)
		return
	}

	if upload.Status == "processed" {
		http.Error(w, "upload already processed", http.StatusBadRequest)
		return
	}

	mockStats := []models.PlayerStats{
		{
			ID:         generateID(),
			GameID:     upload.GameID,
			TeamName:   "Almendra Basketball",
			PlayerName: "Edgar Montenegro",
			Points:     17,
			Rebounds:   1,
			Assists:    1,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         generateID(),
			GameID:     upload.GameID,
			TeamName:   "Almendra Basketball",
			PlayerName: "Nicolás Landoni",
			Points:     18,
			Rebounds:   14,
			Assists:    1,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	playerStats = append(playerStats, mockStats...)

	for i := range uploads {
		if uploads[i].ID == upload.ID {
			uploads[i].Status = "processed"
			uploads[i].ProcessedAt = time.Now()
			break
		}
	}

	response := map[string]any{
		"message":         "analytics processed succesfully",
		"upload_id":       req.UploadID,
		"game_id":         upload.GameID,
		"records_created": len(mockStats),
	}

	writeJSON(w, http.StatusOK, response)
}

func listPlayerStatsByGameID(w http.ResponseWriter, r *http.Request) {
	gameID := strings.TrimPrefix(r.URL.Path, "/analytics/games/")
	gameID = strings.TrimSuffix(gameID, "/players")

	if gameID == "" {
		http.Error(w, "game_id is required", http.StatusBadRequest)
		return
	}

	filteredStats := []models.PlayerStats{}

	for _, stat := range playerStats {
		if stat.GameID == gameID {
			filteredStats = append(filteredStats, stat)
		}
	}

	writeJSON(w, http.StatusOK, filteredStats)
}
