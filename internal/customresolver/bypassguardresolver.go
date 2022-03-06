package customresolver

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/erkexzcx/stoppropaganda/internal/targets"
)

type BypassGuardResolver struct {
}

var errBypassNotFound = errors.New("bypass guard IP not found")

func (cr *BypassGuardResolver) LookupIPAddr(ctx context.Context, host string) (names []net.IPAddr, err error) {
	for _, targetBypass := range targets.BypassIPs {
		if host == targetBypass.Host {
			log.Printf("Using guard bypass %v %v", host, targetBypass.IPs)
			names = targetBypass.IPs
			return
		}
	}
	err = errBypassNotFound
	return
}
