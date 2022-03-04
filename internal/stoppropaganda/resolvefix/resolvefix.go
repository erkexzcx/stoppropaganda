package resolvefix

import (
	"errors"
	"net"
)

func CheckNonPublicIPAddrs(addrs []net.TCPAddr) error {
	ips := make([]net.IP, len(addrs))
	for i, addr := range addrs {
		ips[i] = addr.IP
	}
	return CheckNonPublicIP(ips)
}

func CheckNonPublicIP(ips []net.IP) error {
	if ContainsNonPublicIP(ips) {
		return errors.New("Non public IP detected")
	}
	return nil
}

func ContainsNonPublicIP(ips []net.IP) bool {
	for _, ip := range ips {
		if ip.IsPrivate() || ip.IsLoopback() || ip.IsUnspecified() {
			return true
		}
	}
	return false
}
