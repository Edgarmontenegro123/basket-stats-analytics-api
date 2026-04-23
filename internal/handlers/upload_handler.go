package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/models"
)

var Uploads []models.StatUpload

type createUploadRequest struct {
	GameID   string `json:"game_id"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
}

func UploadsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createUpload(w, r)
	case http.MethodGet:
		listUploads(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func createUpload(w http.ResponseWriter, r *http.Request) {
	var req createUploadRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.GameID == "" {
		http.Error(w, "game_id is required", http.StatusBadRequest)
		return
	}

	if req.FileName == "" {
		http.Error(w, "file_name is required", http.StatusBadRequest)
		return
	}

	if req.FileType == "" {
		http.Error(w, "file_type is required", http.StatusBadRequest)
		return
	}

	upload := models.StatUpload{
		ID:         generateID(),
		GameID:     req.GameID,
		FileName:   req.FileName,
		FileType:   req.FileType,
		Status:     "uploaded",
		UploadedAt: time.Now(),
	}

	Uploads = append(Uploads, upload)

	writeJSON(w, http.StatusCreated, upload)
}

func listUploads(w http.ResponseWriter, _ *http.Request) {
	if Uploads == nil {
		Uploads = []models.StatUpload{}
	}

	writeJSON(w, http.StatusOK, Uploads)
}
