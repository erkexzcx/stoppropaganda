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
	ips := make([]net.IP, len(addrs))
	for i, ia := range addrs {
		ips[i] = ia.IP
	}
	return ips, nil
}

func GetIPs(host string) (ips []net.IP, err error) {

	ipAddresses := make([]net.IP, 0, 1)
	addr := net.ParseIP(host)
	if addr == nil {
		ips, err := CustomLookupIP(host)
		if err != nil {
			return nil, err
		}
		for _, ip := range ips {
			if ipv4 := ip.To4(); ipv4 != nil {
				ipAddresses = append(ipAddresses, ipv4)
			}
		}
	} else {
		ipAddresses = append(ipAddresses, addr)
	}
	return ipAddresses, nil
}
