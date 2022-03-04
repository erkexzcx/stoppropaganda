package resolvefix

import (
	"errors"
	"net"
)

func CheckNonPublicTCPEndpoints(addrs []net.TCPAddr) error {
	ips := make([]net.IP, len(addrs))
	for i, addr := range addrs {
		ips[i] = addr.IP
	}
	return CheckNonPublicIPs(ips)
}

func CheckNonPublicIPs(ips []net.IP) error {
	for _, ip := range ips {
		if IsNonPublic(ip) {
			return errors.New("Non public IP detected: " + ip.String())
		}
	}
	return nil
}

func IsNonPublic(ip net.IP) bool {
	return ip.IsPrivate() || ip.IsLoopback() || ip.IsUnspecified()
}
