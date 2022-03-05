package spdnsclient

import (
	"context"
	"net"
	"sync"
)

// dnsWaitGroup can be used by tests to wait for all DNS goroutines to
// complete. This avoids races on the test hooks.
var dnsWaitGroup sync.WaitGroup

// ipVersion returns the provided network's IP version: '4', '6' or 0
// if network does not end in a '4' or '6' byte.
func ipVersion(network string) byte {
	if network == "" {
		return 0
	}
	n := network[len(network)-1]
	if n != '4' && n != '6' {
		n = 0
	}
	return n
}

// LookupHost looks up the given host using the local resolver.
// It returns a slice of that host's addresses.
func (r *SPResolver) LookupHost(ctx context.Context, host string) (addrs []string, err error) {
	// Make sure that no matter what we do later, host=="" is rejected.
	// parseIP, for example, does accept empty strings.
	if host == "" {
		return nil, &net.DNSError{Err: errNoSuchHost.Error(), Name: host, IsNotFound: true}
	}
	if ip, _ := parseIPZone(host); ip != nil {
		return []string{host}, nil
	}
	return r.lookupHost(ctx, host)
}

func (r *SPResolver) lookupHost(ctx context.Context, host string) (addrs []string, err error) {
	//order := systemConf().hostLookupOrder(r, host)
	// if !r.preferGo() && order == hostLookupCgo {
	// 	if addrs, err, ok := cgoLookupHost(ctx, host); ok {
	// 		return addrs, err
	// 	}
	// 	// cgo not available (or netgo); fall back to Go's DNS resolver
	// 	order = hostLookupFilesDNS
	// }
	return r.GoLookupHost(ctx, host)
}

func (r *SPResolver) lookupIP(ctx context.Context, network, host string) (addrs []net.IPAddr, err error) {
	if true {
		return r.goLookupIP(ctx, network, host)
	}
	// order := systemConf().hostLookupOrder(r, host)
	// if order == hostLookupCgo {
	// 	if addrs, err, ok := cgoLookupIP(ctx, network, host); ok {
	// 		return addrs, err
	// 	}
	// 	// cgo not available (or netgo); fall back to Go's DNS resolver
	// 	order = hostLookupFilesDNS
	// }
	ips, _, err := r.goLookupIPCNAME(ctx, network, host)
	return ips, err
}

func (r *SPResolver) dial(ctx context.Context, network, server string) (net.Conn, error) {
	// Calling Dial here is scary -- we have to be sure not to
	// dial a name that will require a DNS lookup, or Dial will
	// call back here to translate it. The DNS config parser has
	// already checked that all the cfg.servers are IP
	// addresses, which Dial will use without a DNS lookup.
	var c net.Conn
	var err error
	if r != nil && r.Dial != nil {
		c, err = r.Dial(ctx, network, server)
	} else {
		var d net.Dialer
		c, err = d.DialContext(ctx, network, server)
	}
	if err != nil {
		return nil, mapErr(err)
	}
	return c, nil
}
