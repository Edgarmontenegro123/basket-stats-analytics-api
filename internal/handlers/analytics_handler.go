package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/models"
	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/services"
)

var playerStats []models.PlayerStats
var teamStats []models.TeamStat

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

func GameTeamStatsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listTeamStatsByGameID(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func GameStatsRouter(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/players") {
		GamePlayerStatsHandler(w, r)
		return
	}

	if strings.HasSuffix(r.URL.Path, "/teams") {
		GameTeamStatsHandler(w, r)
		return
	}

	http.Error(w, "route not found", http.StatusNotFound)
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

	mockStats, mockTeamStats, updatedUploads, err := services.ProcessAnalytics(upload, uploads, generateID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uploads = updatedUploads

	playerStats = append(playerStats, mockStats...)
	teamStats = append(teamStats, mockTeamStats...)

	response := map[string]any{
		"message":                "analytics processed successfully",
		"upload_id":              req.UploadID,
		"game_id":                upload.GameID,
		"player_records_created": len(mockStats),
		"team_records_created":   len(mockTeamStats),
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

	filteredStats := services.GetPlayerStatsByGameID(playerStats, gameID)

	writeJSON(w, http.StatusOK, filteredStats)
}

func listTeamStatsByGameID(w http.ResponseWriter, r *http.Request) {
	gameID := strings.TrimPrefix(r.URL.Path, "/analytics/games/")
	gameID = strings.TrimSuffix(gameID, "/teams")

	if gameID == "" {
		http.Error(w, "game_id is required", http.StatusBadRequest)
		return
	}

	filteredStats := services.GetTeamStatsByGameID(teamStats, gameID)

	writeJSON(w, http.StatusOK, filteredStats)
}
