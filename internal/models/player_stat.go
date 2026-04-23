package models

import "time"

type PlayerStats struct {
	ID         string    `json:"id"`
	GameID     string    `json:"game_id"`
	TeamName   string    `json:"team_name"`
	PlayerName string    `json:"player_name"`
	Points     int       `json:"points"`
	Rebounds   int       `json:"rebounds"`
	Assists    int       `json:"assists"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
