package customresolver

import (
	"context"
	"net"
	"time"

	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/spdnsclient"
	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/targets"
	"github.com/patrickmn/go-cache"
)

var DnsCache *cache.Cache

func MakeDNSConfig() (conf *spdnsclient.SPDNSConfig) {
	conf = &spdnsclient.SPDNSConfig{
		Ndots:    1,
		Timeout:  5 * time.Second,
		Attempts: 2,
	}
	conf.Servers = targets.ReferenceDNSServersForHTTP

	if len(conf.Search) == 0 {
		conf.Search = spdnsclient.DnsDefaultSearch()
	}
	return
}

var MasterStopPropagandaResolver = &CustomResolver{
	FirstResolver: &spdnsclient.SPResolver{

		CustomDNSConfig: MakeDNSConfig(),
	},
	ParentResolver: net.DefaultResolver,
}

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
	names, err = cr.ParentResolver.LookupIPAddr(ctx, host)
	if err == nil {
		DnsCache.SetDefault(host, names)
		return
	}
	return
}

func init() {
	DnsCache = cache.New(5*time.Minute, 10*time.Minute)
}
