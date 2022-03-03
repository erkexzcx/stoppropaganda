package customresolver

import (
	"context"
	"net"
	"time"

	"github.com/patrickmn/go-cache"
)

var resolver *net.Resolver
var dnscache *cache.Cache

type CustomResolver struct{}

type Resolver interface {
	LookupIPAddr(context.Context, string) (names []net.IPAddr, err error)
}

func (cs *CustomResolver) LookupIPAddr(ctx context.Context, host string) (names []net.IPAddr, err error) {
	if c, found := dnscache.Get(host); found {
		return c.([]net.IPAddr), nil
	}

	names, err = resolver.LookupIPAddr(ctx, host)
	if err == nil {
		dnscache.SetDefault(host, names)
	}
	return
}

func init() {
	resolver = net.DefaultResolver
	dnscache = cache.New(5*time.Minute, 10*time.Minute)
}
