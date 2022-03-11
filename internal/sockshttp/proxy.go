// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package proxy provides support for a variety of protocols to proxy network
// data.
package sockshttp

import (
	"net"
	"net/url"
	"time"
)

// A Dialer is a means to establish a connection.
type Dialer interface {
	// Dial connects to the given address via the proxy.
	Dial(network, addr string) (c net.Conn, err error)
}

// Auth contains authentication parameters that specific Dialers may require.
type Auth struct {
	User, Password string
}

// FromEnvironment returns the dialer specified by the proxy related variables in
// the environment.
func Initialize(proxyParam, proxyBypassParam string, proxyTimeout time.Duration) (ProxyChain, Dialer) {
	if len(proxyParam) == 0 {
		return nil, Direct
	}

	proxyChain := ParseProxyChain(proxyParam)
	proxy := MakeDialerThrough(Direct, proxyChain, proxyTimeout)

	if len(proxyBypassParam) == 0 {
		return proxyChain, proxy
	}

	perHost := NewPerHost(proxy, Direct)
	perHost.AddFromString(proxyBypassParam)
	return proxyChain, perHost
}

// proxySchemes is a map from URL schemes to a function that creates a Dialer
// from a URL with such a scheme.
var proxySchemes map[string]func(*url.URL, Dialer) (Dialer, error)

// RegisterDialerType takes a URL scheme and a function to generate Dialers from
// a URL with that scheme and a forwarding Dialer. Registered schemes are used
// by FromURL.
func RegisterDialerType(scheme string, f func(*url.URL, Dialer) (Dialer, error)) {
	if proxySchemes == nil {
		proxySchemes = make(map[string]func(*url.URL, Dialer) (Dialer, error))
	}
	proxySchemes[scheme] = f
}
