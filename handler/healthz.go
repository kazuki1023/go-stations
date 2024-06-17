package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// `model`パッケージの構造体を使用
	response := model.HealthzResponse{Message: "OK"}

	// レスポンスのContent-Typeを設定
	w.Header().Set("Content-Type", "application/json")

	// Encoder を使用してJSONにシリアライズし、HTTPレスポンスに書き込む
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// エラーハンドリング
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
