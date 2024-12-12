package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/babacar-thiam/go-rbac-api/internal/services"
)

// Handler handles the http requests
type Handler struct {
	service *services.Service
}

// NewHandler creates and return a new Handler instance
func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.service.GetAllRoles()
	if err != nil {
		http.Error(w, "error retrieving roles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(roles); err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
}
