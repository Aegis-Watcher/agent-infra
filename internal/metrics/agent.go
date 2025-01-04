package metrics

import (
	"log"
	"time"

	c "agent-infra/config"
	"agent-infra/internal/utils"
)

type Agent struct {
	config c.Config
}

func NewAgent(config c.Config) *Agent {
	return &Agent{config: config}
}

func (a *Agent) Start() error {
	ticker := time.NewTicker(a.config.CollectionInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Collect and send metrics
			err := a.collectAndSend()
			if err != nil {
				log.Printf("Error sending metrics: %v", err)
			}
		}
	}
}

func (a *Agent) collectAndSend() error {
	// Collect metrics
	cpu, err := GetCPUUsage()
	if err != nil {
		return err
	}

	mem, err := GetMemoryUsage()
	if err != nil {
		return err
	}

	disk, err := GetDiskUsage()
	if err != nil {
		return err
	}

	// Prepare payload
	payload := map[string]interface{}{
		"cpu":         cpu,
		"memory":      mem,
		"disk":        disk,
		"time":        time.Now().UTC(),
		"instance_id": a.config.InstanceID,
	}

	// Compress payload
	compressedPayload, err := utils.CompressJSON(payload)
	if err != nil {
		return err
	}

	// Send payload
	return SendMetrics(a.config.ServerURL, compressedPayload)
}
