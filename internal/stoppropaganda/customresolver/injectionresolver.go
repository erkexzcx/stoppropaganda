package customresolver

import (
	"context"
	"net"
	"time"

	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/targets"
)

var injectionGoResolver = &net.Resolver{
	PreferGo: true,
	Dial: func(ctx context.Context, network, address string) (conn net.Conn, err error) {
		tries := 3
		d := net.Dialer{
			Timeout: time.Millisecond * time.Duration(10000),
		}
		for i := 0; i < tries; i++ {
			// use DNS targets from dns.go
			for dnsTarget, _ := range targets.TargetDNSServers {
				// eg. d.DialContext(ctx, "tcp", "194.54.14.186:53")
				conn, err = d.DialContext(ctx, network, dnsTarget)
				if err == nil {
					// return first working conn to DNS
					return
				}
			}
		}
		return
	},
}
