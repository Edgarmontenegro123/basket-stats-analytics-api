package services

import (
	"errors"
	"time"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/models"
)

var ErrUploadAlreadyProcessed = errors.New("upload already processed")

func GenerateMockAnalytics(gameID string, generateID func() string) ([]models.PlayerStats, []models.TeamStat) {
	now := time.Now()

	playerStats := []models.PlayerStats{
		{
			ID:         generateID(),
			GameID:     gameID,
			TeamName:   "Almendra Basketball",
			PlayerName: "Edgar Montenegro",
			Points:     17,
			Rebounds:   1,
			Assists:    1,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         generateID(),
			GameID:     gameID,
			TeamName:   "Almendra Basketball",
			PlayerName: "Nicolás Landoni",
			Points:     20,
			Rebounds:   1,
			Assists:    1,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	teamStats := []models.TeamStat{
		{
			ID:        generateID(),
			GameID:    gameID,
			TeamName:  "Almendra Basketball",
			Points:    20,
			Rebounds:  1,
			Assists:   1,
			Turnovers: 1,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        generateID(),
			GameID:    gameID,
			TeamName:  "Pegasos",
			Points:    20,
			Rebounds:  1,
			Assists:   1,
			Turnovers: 1,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	return playerStats, teamStats
}

func ValidateUploadStatus(upload models.StatUpload) error {
	if upload.Status == "processed" {
		return ErrUploadAlreadyProcessed
	}

	return nil
}
