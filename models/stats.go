package models

type WebsiteStatus string

const (
	StatusUP   WebsiteStatus = "UP"
	StatusDown WebsiteStatus = "DOWN"
)

type Stats struct {
	Domain         string        `json:"domain"`
	IPv4           string        `json:"ipv4"`
	IPv6           string        `json:"ipv6"`
	StatusCode     int           `json:"status_code"`
	ResponseTimeMs int64         `json:"response_time_ms"`
	Status         WebsiteStatus `json:"status"`
	Error          string        `json:"error,omitempty"`
	DNSError       string        `json:"dns_error,omitempty"`
}
