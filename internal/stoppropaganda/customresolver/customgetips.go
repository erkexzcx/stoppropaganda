package customresolver

import (
	"context"
	"net"
)

// Modified to use stoppropaganda's CustomResolver
// so that it caches DNS records
func CustomLookupIP(host string) ([]net.IP, error) {
	resolver := CustomResolver{
		ParentResolver: net.DefaultResolver,
	}
	addrs, err := resolver.LookupIPAddr(context.Background(), host)
	if err != nil {
		return nil, err
	}
	ips := make([]net.IP, 0, len(addrs))
	for _, ia := range addrs {
		ips = append(ips, ia.IP)
	}
	return ips, nil
}

func GetIPs(host string) (ips []net.IP, err error) {
	addr := net.ParseIP(host)
	if addr == nil {
		ips, err := CustomLookupIP(host)
		if err != nil {
			return nil, err
		}
		ipAddresses := make([]net.IP, 0, len(ips))
		for _, ip := range ips {
			if ipv4 := ip.To4(); ipv4 != nil {
				ipAddresses = append(ipAddresses, ipv4)
			}
		}
		return ipAddresses, nil
	}
	return []net.IP{addr}, nil
}
