package customresolver

import (
	"context"
	"net"
	"time"

	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/targets"
)

var InjectionGoResolver = &net.Resolver{
	PreferGo: true,
	Dial: func(ctx context.Context, network, address string) (conn net.Conn, err error) {
		d := net.Dialer{
			Timeout: time.Millisecond * time.Duration(1000),
		}

		// use DNS targets from dns.go
		for _, dnsTarget := range targets.ReferenceDNSServersForHTTP {
			// eg. d.DialContext(ctx, "tcp", "194.54.14.186:53")
			conn, err = d.DialContext(ctx, network, dnsTarget)
			if err == nil {
				// return first working conn to DNS
				return
			}
		}

		return d.DialContext(ctx, network, address)
	},
}
