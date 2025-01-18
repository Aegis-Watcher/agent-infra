package metrics

import "os"

// GetHostname retrieves the current machine's hostname
func GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostname, nil
}
