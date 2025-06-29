package handlers

import (
	"awesomeProject10/internal/config"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupHandler() *Handler {
	cfg := &config.Config{}
	cfg.Sizes = []int{250, 500, 1000, 2000, 5000}
	return NewHandler(cfg)
}

func TestOptimize_ValidRequest(t *testing.T) {
	h := setupHandler()

	body := `{"items": 12001}`
	req := httptest.NewRequest(http.MethodPost, "/calculate", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h.Optimize(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	respBody, _ := io.ReadAll(resp.Body)

	var parsed struct {
		Packs map[int]int `json:"packs"`
		Total int         `json:"total"`
	}
	err := json.Unmarshal(respBody, &parsed)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, parsed.Total, 12001)
	assert.NotEmpty(t, parsed.Packs)
}

func TestOptimize_InvalidJSON(t *testing.T) {
	h := setupHandler()

	body := `{"items": "abc"}`
	req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h.Optimize(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestOptimize_NegativeValue(t *testing.T) {
	h := setupHandler()

	body := `{"items": -10}`
	req := httptest.NewRequest(http.MethodPost, "/calculate", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h.Optimize(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestOptimize_UnknownField(t *testing.T) {
	h := setupHandler()

	body := `{"itemz": 100}`
	req := httptest.NewRequest(http.MethodPost, "/calculate", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h.Optimize(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
