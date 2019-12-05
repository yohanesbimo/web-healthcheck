package api

import (
	"net/http"
	"time"
)

// Healthchecker Check URL Health. Return false if not healthy and true if healthy
func Healthchecker(url string) bool {
	client := http.Client{
		Timeout: 800 * time.Millisecond,
	}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
