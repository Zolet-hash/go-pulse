package dns

import (
	"context"
	"net"
	"time"
)

// Resolve hostname into IPv4 and Ipv6 addresses.
func ResolveIPAddresses(hostname string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resolver := net.DefaultResolver

	addresses, err := resolver.LookupIPAddr(ctx, hostname)
	if err != nil {
		return "", "", err
	}

	var ipv4 string
	var ipv6 string

	for _, address := range addresses {
		if address.IP.To4() != nil {
			if ipv4 == "" {
				ipv4 = address.IP.String()
			}
		} else if address.IP.To16() != nil {
			if ipv6 == "" {
				ipv6 = address.IP.String()
			}
		}
	}
	return ipv4, ipv6, nil
}
