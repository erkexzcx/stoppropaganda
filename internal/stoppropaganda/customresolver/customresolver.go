package customresolver

import (
	"context"
	"net"
	"time"

	"github.com/patrickmn/go-cache"
)

var dnscache *cache.Cache

type CustomResolver struct {
	ParentResolver *net.Resolver
}

type Resolver interface {
	LookupIPAddr(context.Context, string) (names []net.IPAddr, err error)
}

func (cr *CustomResolver) LookupIPAddr(ctx context.Context, host string) (names []net.IPAddr, err error) {
	if c, found := dnscache.Get(host); found {
		return c.([]net.IPAddr), nil
	}

	names, err = cr.ParentResolver.LookupIPAddr(ctx, host)
	if err == nil {
		dnscache.SetDefault(host, names)
	}
	return
}

func init() {
	dnscache = cache.New(5*time.Minute, 10*time.Minute)
}
