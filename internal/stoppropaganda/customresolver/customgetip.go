package customresolver

import (
	"context"
	"net"
)

// Modified to use stoppropaganda's CustomResolver
// so that it caches DNS records
func CustomLookupIP(host string, helperIPBuf []net.IP) ([]net.IP, error) {
	addrs, err := MasterStopPropagandaResolver.LookupIPAddr(context.Background(), host)
	if err != nil {
		return nil, err
	}
	ips := helperIPBuf[:0]
	for _, ia := range addrs {
		ips = append(ips, ia.IP)
	}
	return ips, nil
}

func GetIPs(host string, helperIPBuf []net.IP) (ips []net.IP, err error) {
	addr := net.ParseIP(host)
	if addr == nil {
		ips, err := CustomLookupIP(host, helperIPBuf)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(ips); i++ {
			ip := ips[i]
			if ipv4 := ip.To4(); ipv4 == nil {
				// Remove inplace trick
				// - swap i and last element
				ips[i] = ips[len(ips)-1]
				// - pop last element
				ips[len(ips)-1] = nil
				ips = ips[:len(ips)-1]
			}
		}
		return ips, nil
	}
	helperIPBuf[0] = addr
	return helperIPBuf[:1], nil
}
