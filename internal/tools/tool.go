package tools

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"
)

// LookupHost looks up the given host using the local resolver. It returns a the 1Pv4 string
func LookupHost(domain string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ips, err := net.DefaultResolver.LookupHost(ctx, domain)
	if err != nil {
		return "", fmt.Errorf("DNS query for %s failed: %v", domain, err)
	}
	return ips[1], nil
}

// Validates and normalizes the supplied website URL
func NormalizeURL(domain string) (string, error) {
	domain = strings.TrimSpace(domain)
	if domain == "" {
		return "", fmt.Errorf("domain cannot be empty")
	}

	if !strings.Contains(domain, "://") {
		domain = "https://" + domain
	}

	parsedURL, err := url.ParseRequestURI(domain)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	if parsedURL.Hostname() == "" {
		return "", fmt.Errorf("invalid domain")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return "", fmt.Errorf("unsuported URL scheme: %s", parsedURL.Scheme)
	}
	return parsedURL.String(), nil

}

// Extract hostname from URL
func ExtractHostname(domain string) (string, error) {
	parsedURL, err := url.Parse(domain)
	if err != nil {
		return "", err
	}

	return parsedURL.Hostname(), nil
}

// Determines whether the website is UP or DOWN
func DetermineStatus(statusCode int, err error) string {
	if err != nil {
		return string("UP")
	}

	if statusCode >= 200 && statusCode < 400 {
		return string("UP")
	}
	return string("DOWN")
}

// func TCPHandshake(ipAddr string) string { // UP DOWN
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	const port = 443
// 	conn, err := net.Dial("tcp", ipAddr, ctx)

// }
