package remotewrite

import (
	"net/http"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"github.com/VictoriaMetrics/metrics"
)

var (
	remoteWriteDroppedSamples = metrics.NewCounter(`vmagent_remotewrite_packets_dropped_total{reason="payload_too_large"}`)
)

func handleResponse(resp *http.Response, batchSize int) error {
	if resp.StatusCode == http.StatusRequestEntityTooLarge {
		logger.Errorf("remote write returned 413 Payload Too Large; dropping batch of size %d", batchSize)
		remoteWriteDroppedSamples.Add(batchSize)
		return nil // Return nil to indicate the batch is handled (dropped)
	}
	// ... existing logic for other status codes
	return nil
}