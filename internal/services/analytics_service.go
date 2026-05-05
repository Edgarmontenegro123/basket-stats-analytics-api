package services

import (
	"errors"
	"os"
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

func ProcessAnalytics(upload models.StatUpload, uploads []models.StatUpload, generateID func() string) ([]models.PlayerStats, []models.TeamStat, []models.StatUpload, error) {
	err := ValidateUploadStatus(upload)
	if err != nil {
		return nil, nil, uploads, err
	}

	if upload.FilePath == "" {
		return nil, nil, uploads, errors.New("file path is required")
	}

	file, err := os.Open(upload.FilePath)
	if err != nil {
		return nil, nil, uploads, errors.New("uploaded file not found")
	}

	defer func() {
		_ = file.Close()
	}()

	pdfText, err := ExtractTextFromPDF(upload.FilePath)
	if err != nil {
		return nil, nil, uploads, errors.New("error extracting text from pdf")
	}

	if pdfText == "" {
		return nil, nil, uploads, errors.New("pdf text is empty")
	}

	playerStats, teamStats := GenerateMockAnalytics(upload.GameID, generateID)

	for i := range uploads {
		if uploads[i].ID == upload.ID {
			uploads[i].Status = "processed"
			uploads[i].ProcessedAt = time.Now()
			break
		}
	}

	return playerStats, teamStats, uploads, nil
}

func GetPlayerStatsByGameID(stats []models.PlayerStats, gameID string) []models.PlayerStats {
	var filteredStats []models.PlayerStats

	for _, stat := range stats {
		if stat.GameID == gameID {
			filteredStats = append(filteredStats, stat)
		}
	}

	return filteredStats
}

func GetTeamStatsByGameID(stats []models.TeamStat, gameID string) []models.TeamStat {
	var filteredStats []models.TeamStat

	for _, stat := range stats {
		if stat.GameID == gameID {
			filteredStats = append(filteredStats, stat)
		}
	}

	return filteredStats
}
