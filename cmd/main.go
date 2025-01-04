package main

import (
	"fmt"
	"log"
	"os"
	"time"

	c "agent-infra/config"
	"agent-infra/internal/metrics"
	"agent-infra/internal/utils"
)

func main() {
	var serverURL string

	ENV := os.Getenv("ENV")

	if ENV == "" {
		ENV = "production"
	}

	if ENV == "production" {
		fmt.Println("Running in production environment")
		serverURL = "https://api.aegiswatcher.com/api/v1/metrics"
	} else {
		fmt.Println("Running in development environment")
		serverURL = "http://localhost:8080/api/v1/metrics"
	}

	instanceID := os.Getenv("AEGIS_WATCHER_METRICS_AGENT_INSTANCE_ID")
	if instanceID == "" {
		// generate a instanceID if not provided in the environment variables
		instanceID = utils.GenerateInstanceID()
	}

	// Initialize configuration (e.g., collection intervals)
	config := c.Config{
		ServerURL:          serverURL,
		CollectionInterval: 10 * time.Second, // Collect every 10 seconds
		InstanceID:         os.Getenv("AEGIS_WATCHER_METRICS_AGENT_INSTANCE_ID"),
	}
	fmt.Println("config: ", config)

	// Start metric collection
	agent := metrics.NewAgent(config)
	err := agent.Start()
	if err != nil {
		log.Fatalf("Error starting agent: %v", err)
	}
}
