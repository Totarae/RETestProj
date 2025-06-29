package handlers

import (
	"awesomeProject10/internal/config"
	"awesomeProject10/internal/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	config *config.Config
}

// TODO: replace packsize with config
// default
func NewHandler(conf *config.Config) *Handler {
	return &Handler{config: conf}
}

// ServeIndex — UI method
func (h *Handler) ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}

// Optimize — main API method
func (h *Handler) Optimize(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Items int `json:"items"`
	}
	type response struct {
		Packs map[int]int `json:"packs"`
		Total int         `json:"total"`
	}

	var req request
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid JSON or type mismatch", http.StatusBadRequest)
		return
	}

	if req.Items <= 0 {
		http.Error(w, "items must be a positive integer", http.StatusBadRequest)
		return
	}

	packSizes := h.config.GetPackSizes()

	result := service.OptimizePacks(req.Items, packSizes)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{
		Packs: result.Packs,
		Total: result.Total,
	})
}
