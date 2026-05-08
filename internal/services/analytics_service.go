package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/models"
)

func ValidateUploadStatus(upload models.StatUpload) error {
	if upload.Status == "processed" {
		return errors.New("upload already processed")
	}

	if upload.Status != "uploaded" {
		return errors.New("upload is not ready to be processed")
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

	teamStats, err := ParseTeamStatsFromText(pdfText, upload.GameID, generateID)
	if err != nil {
		return nil, nil, uploads, err
	}

	playerStats := ParsePlayerStatsFromText(pdfText, upload.GameID)

	for i := range playerStats {
		playerStats[i].ID = fmt.Sprintf("%s-%d", generateID(), i+1)
	}

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
