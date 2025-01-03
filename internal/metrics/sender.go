package metrics

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

func SendMetrics(url string, compressedData []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(compressedData))
	if err != nil {
		return err
	}

	agentToken := AgentToken()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")
	req.Header.Set("Agent-Token", agentToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send metrics, status code: %d", resp.StatusCode)
	}

	fmt.Println("Metrics sent successfully")

	return nil
}

func AgentToken() string {
	agentUUID := os.Getenv("AGENT_UUID")
	agentSecret := os.Getenv("AGENT_SECRET")

	// encode the agentUUID and agentSecret to base64 e.g "agentUUID:agentSecret" and return as agentToken

	agentToken := base64.StdEncoding.EncodeToString([]byte(agentUUID + ":" + agentSecret))

	return agentToken
}
