package httpcheck

import (
	"net/http"
	"time"
)

// Sends an HTTP message and measures how“ long it takes
func CHeckHTTPStatus(client *http.Client, domain string) (int, int64, error) {
	start := time.Now()
	req, err := http.NewRequest(
		http.MethodGet, domain, nil,
	)
	if err != nil {
		return 0, 0, err
	}

	req.Header.Set("User-Agent", "Go-website-Monitor/1.0")
	resp, err := client.Do(req)
	responseTime := time.Since(start).Milliseconds()

	if err != nil {
		return 0, responseTime, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, responseTime, nil
}


const requestTimeout = 10 * time.Second

// Create and configure an HTTP client
func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: requestTimeout,

		// Prevent redirects from hiding the original status code
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

}