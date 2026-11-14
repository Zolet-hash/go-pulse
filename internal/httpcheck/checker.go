package httpcheck

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	requestTimeout       = 10 * time.Second
	maxResponseDrainSize = 64 * 1024
)

// Create and configure an HTTP client
func NewHTTPClient() *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,

		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,

		ForceAttemptHTTP2:   true,
		MaxIdleConns:        500,
		MaxIdleConnsPerHost: 4,
		MaxConnsPerHost:     4,

		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 8 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   requestTimeout,

		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

}

// Sends an HTTP message and measures how“ long it takes
func CHeckHTTPStatus(ctx context.Context, client *http.Client, domain string) (int, int64, error) {
	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet, domain, nil,
	)

	if err != nil {
		return 0, 0, err
	}

	req.Header.Set("User-Agent", "Go-website-Monitor/1.0")
	start := time.Now()

	resp, err := client.Do(req)
	responseTime := time.Since(start).Milliseconds()

	if err != nil {
		return 0, responseTime, err
	}
	defer resp.Body.Close()

	// Drain a bounded amount to improve connection reuse.
	// Never consume an unlimited response body
	_, _ = io.Copy(
		io.Discard,
		io.LimitReader(resp.Body, maxResponseDrainSize),
	)

	return resp.StatusCode, responseTime, nil
}
