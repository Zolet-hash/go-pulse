package monitor

import (
	"github.com/Zolet-hash/go-pulse/internal/dns"
	"github.com/Zolet-hash/go-pulse/internal/httpcheck"
	"github.com/Zolet-hash/go-pulse/internal/tools"
	"github.com/Zolet-hash/go-pulse/models"
)

// Ochestrate the individual operations
func MonitorWebsite(domain string) models.Stats {
	stats := models.Stats{
		Domain: domain,
		Status: "DOWN",
	}

	normalizedURL, err := tools.NormalizeURL(domain)
	if err != nil {
		return stats
	}

	stats.Domain = normalizedURL

	hostname, err := tools.ExtractHostname(normalizedURL)
	if err != nil {
		return stats
	}

	ipv4, ipv6, dnsErr := dns.ResolveIPAddresses(hostname)

	stats.IPv4 = ipv4
	stats.IPv6 = ipv6

	client := httpcheck.NewHTTPClient()

	statusCode, responseTime, httpErr := httpcheck.CHeckHTTPStatus(client, normalizedURL)

	stats.StatusCode = statusCode
	stats.ResponseTimeMs = responseTime
	stats.Status = tools.DetermineStatus(statusCode, httpErr)

	//DNS failure -> we couldn't resolve the website
	if dnsErr != nil {
		stats.Status = "DOWN"
	}

	return stats
}
