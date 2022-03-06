package targets

import "net"

type BypassWebsite struct {
	Host string
	IPs  []net.IPAddr
}

var BypassIPs = []BypassWebsite{
	// Looks like they DROP now with iptables
	/*
		{
			Host: "www.rt.com", IPs: []net.IPAddr{
				*mustResolveIPAddr("207.244.80.179"),
				*mustResolveIPAddr("37.48.108.112"),
			},
		},
	*/
}

func mustResolveIPAddr(address string) (ip *net.IPAddr) {
	ip, err := net.ResolveIPAddr("ip", address)
	if err != nil {
		panic(err)
	}
	return
}
