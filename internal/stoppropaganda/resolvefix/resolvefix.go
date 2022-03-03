package resolvefix

import (
	"errors"
	"net"
)

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
