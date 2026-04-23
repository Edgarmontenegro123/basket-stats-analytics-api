package models

import "time"

type StatUpload struct {
	ID          string    `json:"id"`
	GameID      string    `json:"game_id"`
	FileName    string    `json:"file_name"`
	FileType    string    `json:"file_type"`
	Status      string    `json:"status"`
	UploadedAt  time.Time `json:"uploaded_at"`
	ProcessedAt time.Time `json:"processed_at,omitempty"`
}
