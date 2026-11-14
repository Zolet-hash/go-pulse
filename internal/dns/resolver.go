package dns

import (
	"context"
	"net"
)

// ResolveIPAddresses resolves the first available IPv4 & IPv6 addresses
func ResolveIPAddresses(ctx context.Context, hostname string) (string, string, error) {

	addresses, err := net.DefaultResolver.LookupIPAddr(ctx, hostname)
	if err != nil {
		return "", "", err
	}

	var ipv4, ipv6 string

	for _, address := range addresses {
		switch {
		case address.IP.To4() != nil && ipv4 == "":
			ipv4 = address.IP.String()

		case address.IP.To16() != nil && ipv6 == "":
			ipv6 = address.IP.String()
		}
	}
	return ipv4, ipv6, nil
}
