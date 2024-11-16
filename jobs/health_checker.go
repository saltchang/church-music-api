// jobs/health_checker.go

package jobs

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
}

type HealthChecker struct {
	targetURL string
	interval  time.Duration
	client    *http.Client
}

func NewHealthChecker(targetURL string, interval time.Duration) *HealthChecker {
	return &HealthChecker{
		targetURL: targetURL,
		interval:  interval,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (h *HealthChecker) checkHealth() error {
	resp, err := h.client.Get(h.targetURL)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	var healthResponse HealthResponse
	if err := json.Unmarshal(body, &healthResponse); err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}

	log.Printf("Health check status: %s", healthResponse.Status)
	return nil
}

func (h *HealthChecker) Start() {
	ticker := time.NewTicker(h.interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := h.checkHealth(); err != nil {
					log.Printf("Health check failed: %v", err)
				}
			}
		}
	}()
}
