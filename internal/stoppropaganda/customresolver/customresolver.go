package customresolver

import (
	"context"
	"net"
	"time"

	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/spdnsclient"
	"github.com/patrickmn/go-cache"
)

var DnsCache *cache.Cache

type CustomResolver struct {
	FirstResolver  *spdnsclient.SPResolver
	ParentResolver *net.Resolver
}
type Resolver interface {
	LookupIPAddr(context.Context, string) (names []net.IPAddr, err error)
}

func (cr *CustomResolver) LookupIPAddr(ctx context.Context, host string) (names []net.IPAddr, err error) {
	if c, found := DnsCache.Get(host); found {
		return c.([]net.IPAddr), nil
	}
	names, err = cr.FirstResolver.LookupIPAddr(ctx, host)
	if err == nil {
		DnsCache.SetDefault(host, names)
		return
	}
	// names, err = cr.ParentResolver.LookupIPAddr(ctx, host)
	// if err == nil {
	// 	DnsCache.SetDefault(host, names)
	// 	return
	// }
	return
}

func init() {
	DnsCache = cache.New(5*time.Minute, 10*time.Minute)
}
