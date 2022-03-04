package resolvefix

import (
	"errors"
	"net"
)

func CheckNonPublicTCPEndpoints(addrs []net.TCPAddr) error {
	ips := make([]net.IP, 0, len(addrs))
	for _, addr := range addrs {
		ips = append(ips, addr.IP)
	}
	return CheckNonPublicIPs(ips)
}

func CheckNonPublicIPs(ips []net.IP) error {
	for _, ip := range ips {
		if ip.IsPrivate() || ip.IsLoopback() || ip.IsUnspecified() {
			return errors.New("Non public IP detected: " + ip.String())
		}
	}
	return nil
}
