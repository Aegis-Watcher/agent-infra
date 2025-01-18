package metrics

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// SendMetrics sends compressed metrics data to the given URL
func SendMetrics(url string, compressedData []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(compressedData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	agentToken := AgentToken()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")
	req.Header.Set("Agent-Token", agentToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read the response body for detailed error message
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			return fmt.Errorf("failed to send metrics, status code: %d, error reading response: %v", resp.StatusCode, readErr)
		}
		return fmt.Errorf("failed to send metrics, status code: %d, error message: %s", resp.StatusCode, string(body))
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
