package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/models"
)

var uploads []models.StatUpload

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

func UploadByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUploadByID(w, r)
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

	if req.FileType != "pdf" {
		http.Error(w, "file_type must be pdf", http.StatusBadRequest)
		return
	}

	for _, u := range uploads {
		if u.GameID == req.GameID {
			http.Error(w, "upload already exists for this game_id", http.StatusBadRequest)
			return
		}
	}

	upload := models.StatUpload{
		ID:         generateID(),
		GameID:     req.GameID,
		FileName:   req.FileName,
		FileType:   req.FileType,
		Status:     "uploaded",
		UploadedAt: time.Now(),
	}

	uploads = append(uploads, upload)

	writeJSON(w, http.StatusCreated, upload)
}

func listUploads(w http.ResponseWriter, _ *http.Request) {
	if uploads == nil {
		uploads = []models.StatUpload{}
	}

	writeJSON(w, http.StatusOK, uploads)
}

func getUploadByID(w http.ResponseWriter, r *http.Request) {
	uploadID := strings.TrimPrefix(r.URL.Path, "/uploads/")
	if uploadID == "" {
		http.Error(w, "upload id is required", http.StatusBadRequest)
		return
	}

	upload, found := findUploadByID(uploadID)
	if !found {
		http.Error(w, "upload not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, upload)
}

func findUploadByID(uploadID string) (models.StatUpload, bool) {
	for _, upload := range uploads {
		if upload.ID == uploadID {
			return upload, true
		}
	}

	return models.StatUpload{}, false
}
