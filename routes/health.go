package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Uptime  string `json:"uptime"`
	Version string `json:"version"`
}

var start = time.Now()

func formatUptime(d time.Duration) string {
	days := d / (24 * time.Hour)
	d %= 24 * time.Hour
	hours := d / time.Hour
	d %= time.Hour
	minutes := d / time.Minute
	d %= time.Minute
	seconds := d / time.Second

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
	} else if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}
	return fmt.Sprintf("%ds", seconds)
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	healthy := true

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	resp := HealthResponse{
		Status:  "ok",
		Uptime:  formatUptime(time.Since(start)),
		Version: "1.0.0",
	}

	if !healthy {
		w.WriteHeader(http.StatusServiceUnavailable) // 503
		resp.Status = "unhealthy"
	} else {
		w.WriteHeader(http.StatusOK) // 200
	}

	json.NewEncoder(w).Encode(resp)
}
