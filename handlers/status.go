package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/zedithx/devops_practice/internal/metrics"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	metrics.StatusHits.Inc() //increment counter
	status := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}