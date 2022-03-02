package stoppropaganda

import (
	"math/rand"
	"net"
)

var randomDomainRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func getRandomDomain() string {
	randomLength := rand.Intn(20-6) + 6 // from 6 to 20 characters length + ".ru"
	b := make([]rune, randomLength)
	for i := range b {
		b[i] = randomDomainRunes[rand.Intn(len(randomDomainRunes))]
	}
	return string(b) + ".ru"
}

func getIPs(host string) (ips []net.IP, err error) {
	ipAddresses := make([]net.IP, 0, 1)
	addr := net.ParseIP(host)
	if addr == nil {
		ips, err := net.LookupIP(host)
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

func containsPrivateIP(ips []net.IP) bool {
	for _, ip := range ips {
		if ip.IsPrivate() || ip.IsLoopback() {
			return true
		}
	}
	return false
}
