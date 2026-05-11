package services

import (
	"fmt"
	"net/http"
)

func ValidateGameExists(gameID string) error {
	if gameID == "" {
		return fmt.Errorf("game_id is required")
	}

	url := fmt.Sprintf("http://localhost:8080/games/%s", gameID)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("management api is not available")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("game not found")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error validating game")
	}

	return nil
}
