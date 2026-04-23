package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/models"
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

	mockTeamStats := []models.TeamStat{
		{
			ID:        generateID(),
			GameID:    upload.GameID,
			TeamName:  "Almendra Basketball",
			Points:    35,
			Rebounds:  51,
			Assists:   4,
			Turnovers: 17,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        generateID(),
			GameID:    upload.GameID,
			TeamName:  "Pegasos",
			Points:    47,
			Rebounds:  46,
			Assists:   9,
			Turnovers: 12,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	playerStats = append(playerStats, mockStats...)
	teamStats = append(teamStats, mockTeamStats...)

	for i := range uploads {
		if uploads[i].ID == upload.ID {
			uploads[i].Status = "processed"
			uploads[i].ProcessedAt = time.Now()
			break
		}
	}

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

	var filteredStats []models.PlayerStats

	for _, stat := range playerStats {
		if stat.GameID == gameID {
			filteredStats = append(filteredStats, stat)
		}
	}

	writeJSON(w, http.StatusOK, filteredStats)
}

func listTeamStatsByGameID(w http.ResponseWriter, r *http.Request) {
	gameID := strings.TrimPrefix(r.URL.Path, "/analytics/games/")
	gameID = strings.TrimSuffix(gameID, "/teams")

	if gameID == "" {
		http.Error(w, "game_id is required", http.StatusBadRequest)
		return
	}

	var filteredStats []models.TeamStat

	for _, stat := range teamStats {
		if stat.GameID == gameID {
			filteredStats = append(filteredStats, stat)
		}
	}

	writeJSON(w, http.StatusOK, filteredStats)
}
