package crawler

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 30 * time.Second, // Set a timeout
}

func (c *Crawler) fetch(url string) (*http.Response, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// Set headers: e.g., User-Agent
	req.Header.Set("User-Agent", "our-crawler-name")
	// Use the client to send the request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the response
	return res, nil
}
