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

func isPrivateIP(domain string) (ipIsPrivate bool, err error) {
	ipAddresses := make([]net.IP, 0, 1)
	addr := net.ParseIP(domain)
	if addr == nil {
		ips, err := net.LookupIP(domain)
		if err != nil {
			return false, err
		}
		for _, ip := range ips {
			if ipv4 := ip.To4(); ipv4 != nil {
				ipAddresses = append(ipAddresses, ipv4)
			}
		}
	} else {
		ipAddresses = append(ipAddresses, addr)
	}

	for _, ip := range ipAddresses {
		if ip.IsPrivate() || ip.IsLoopback() {
			return true, nil
		}
	}
	return false, nil
}
