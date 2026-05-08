package services

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/models"
	"github.com/ledongthuc/pdf"
)

func ExtractTextFromPDF(filePath string) (string, error) {
	file, reader, err := pdf.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	var buffer bytes.Buffer

	totalPages := reader.NumPage()
	for pageIndex := 1; pageIndex <= totalPages; pageIndex++ {
		page := reader.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			return "", err
		}

		buffer.WriteString(text)
	}

	return buffer.String(), nil
}

func ParseTeamStatsFromText(text string, gameID string, generateID func() string) ([]models.TeamStat, error) {

	lines := strings.Split(text, "\n")

	var totals [][]string

	for i, line := range lines {
		if strings.TrimSpace(line) == "Total" {
			if i+1 >= len(lines) {
				continue
			}

			nextValue := strings.TrimSpace(lines[i+1])
			if _, err := strconv.Atoi(nextValue); err != nil {
				continue
			}

			var values []string

			for j := i + 1; j < len(lines) && len(values) < 27; j++ {
				value := strings.TrimSpace(lines[j])
				if value == "" {
					continue
				}

				values = append(values, value)
			}

			totals = append(totals, values)
		}
	}

	if len(totals) < 2 {
		return nil, errors.New("team stats not found")
	}

	teamNames := []string{"Almendra Basketball", "Pegasos"}
	var teamStats []models.TeamStat

	for i, values := range totals[:2] {
		points, err := strconv.Atoi(values[0])
		if err != nil {
			return nil, err
		}

		rebounds, err := strconv.Atoi(values[17])
		if err != nil {
			return nil, err
		}

		assists, err := strconv.Atoi(values[18])
		if err != nil {
			return nil, err
		}

		turnovers, err := strconv.Atoi(values[19])
		if err != nil {
			return nil, err
		}

		now := time.Now()

		teamStats = append(teamStats, models.TeamStat{
			ID:        generateID(),
			GameID:    gameID,
			TeamName:  teamNames[i],
			Points:    points,
			Rebounds:  rebounds,
			Assists:   assists,
			Turnovers: turnovers,
			CreatedAt: now,
			UpdatedAt: now,
		})
	}

	return teamStats, nil
}
