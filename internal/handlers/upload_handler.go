package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Edgarmontenegro123/basket-stats-analytics-api/internal/models"
)

var uploads []models.StatUpload

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
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "invalid multipart form", http.StatusBadRequest)
		return
	}

	gameID := r.FormValue("game_id")
	if gameID == "" {
		http.Error(w, "game_id is required", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file is required", http.StatusBadRequest)
		return
	}

	defer func() {
		_ = file.Close()
	}()

	if fileHeader.Filename == "" {
		http.Error(w, "file_name is required", http.StatusBadRequest)
		return
	}

	if !strings.HasSuffix(strings.ToLower(fileHeader.Filename), ".pdf") {
		http.Error(w, "file must be a pdf", http.StatusBadRequest)
		return
	}

	for _, u := range uploads {
		if u.GameID == gameID {
			http.Error(w, "upload already exists for this game_id", http.StatusBadRequest)
			return
		}
	}

	upload := models.StatUpload{
		ID:         generateID(),
		GameID:     gameID,
		FileName:   fileHeader.Filename,
		FileType:   "pdf",
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
