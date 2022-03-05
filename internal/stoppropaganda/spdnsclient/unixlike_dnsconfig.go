// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

// Read system DNS config from /etc/resolv.conf

package spdnsclient

import (
	"os"
	"strings"
	"sync/atomic"
	"time"
)

var (
	defaultNS   = []string{"127.0.0.1:53", "[::1]:53"}
	getHostname = os.Hostname // variable for testing
)

type SPDNSConfig struct {
	Servers       []string      // server addresses (in host:port form) to use
	Search        []string      // rooted suffixes to append to local name
	Ndots         int           // number of dots in name to trigger absolute lookup
	Timeout       time.Duration // wait before giving up on a query, including retries
	Attempts      int           // lost packets before giving up on server
	Rotate        bool          // round robin among servers
	UnknownOpt    bool          // anything unknown was encountered
	Lookup        []string      // OpenBSD top-level database "lookup" order
	Err           error         // any error that occurs during open of resolv.conf
	Mtime         time.Time     // time of resolv.conf modification
	Soffset       uint32        // used by serverOffset
	SingleRequest bool          // use sequential A and AAAA queries instead of parallel queries
	UseTCP        bool          // force usage of TCP for DNS resolutions
}

// See resolv.conf(5) on a Linux machine.
/*
func dnsReadConfig(filename string) *SPDNSConfig {
	conf := &SPDNSConfig{
		ndots:    1,
		timeout:  5 * time.Second,
		attempts: 2,
	}
	file, err := os.Open(filename)
	if err != nil {
		conf.servers = defaultNS
		conf.search = dnsDefaultSearch()
		conf.err = err
		return conf
	}
	defer file.Close()
	if fi, err := file.Stat(); err == nil {
		conf.mtime = fi.ModTime()
	} else {
		conf.servers = defaultNS
		conf.search = dnsDefaultSearch()
		conf.err = err
		return conf
	}
	for line, ok := file.ReadLine(); ok; line, ok = file.readLine() {
		if len(line) > 0 && (line[0] == ';' || line[0] == '#') {
			// comment.
			continue
		}
		f := getFields(line)
		if len(f) < 1 {
			continue
		}
		switch f[0] {
		case "nameserver": // add one name server
			if len(f) > 1 && len(conf.servers) < 3 { // small, but the standard limit
				// One more check: make sure server name is
				// just an IP address. Otherwise we need DNS
				// to look it up.
				if parseIPv4(f[1]) != nil {
					conf.servers = append(conf.servers, JoinHostPort(f[1], "53"))
				} else if ip, _ := parseIPv6Zone(f[1]); ip != nil {
					conf.servers = append(conf.servers, JoinHostPort(f[1], "53"))
				}
			}

		case "domain": // set search path to just this domain
			if len(f) > 1 {
				conf.search = []string{ensureRooted(f[1])}
			}

		case "search": // set search path to given servers
			conf.search = make([]string, len(f)-1)
			for i := 0; i < len(conf.search); i++ {
				conf.search[i] = ensureRooted(f[i+1])
			}

		case "options": // magic options
			for _, s := range f[1:] {
				switch {
				case hasPrefix(s, "ndots:"):
					n, _, _ := dtoi(s[6:])
					if n < 0 {
						n = 0
					} else if n > 15 {
						n = 15
					}
					conf.ndots = n
				case hasPrefix(s, "timeout:"):
					n, _, _ := dtoi(s[8:])
					if n < 1 {
						n = 1
					}
					conf.timeout = time.Duration(n) * time.Second
				case hasPrefix(s, "attempts:"):
					n, _, _ := dtoi(s[9:])
					if n < 1 {
						n = 1
					}
					conf.attempts = n
				case s == "rotate":
					conf.rotate = true
				case s == "single-request" || s == "single-request-reopen":
					// Linux option:
					// http://man7.org/linux/man-pages/man5/resolv.conf.5.html
					// "By default, glibc performs IPv4 and IPv6 lookups in parallel [...]
					//  This option disables the behavior and makes glibc
					//  perform the IPv6 and IPv4 requests sequentially."
					conf.singleRequest = true
				case s == "use-vc" || s == "usevc" || s == "tcp":
					// Linux (use-vc), FreeBSD (usevc) and OpenBSD (tcp) option:
					// http://man7.org/linux/man-pages/man5/resolv.conf.5.html
					// "Sets RES_USEVC in _res.options.
					//  This option forces the use of TCP for DNS resolutions."
					// https://www.freebsd.org/cgi/man.cgi?query=resolv.conf&sektion=5&manpath=freebsd-release-ports
					// https://man.openbsd.org/resolv.conf.5
					conf.useTCP = true
				default:
					conf.unknownOpt = true
				}
			}

		case "lookup":
			// OpenBSD option:
			// https://www.openbsd.org/cgi-bin/man.cgi/OpenBSD-current/man5/resolv.conf.5
			// "the legal space-separated values are: bind, file, yp"
			conf.lookup = f[1:]

		default:
			conf.unknownOpt = true
		}
	}
	if len(conf.servers) == 0 {
		conf.servers = defaultNS
	}
	if len(conf.search) == 0 {
		conf.search = dnsDefaultSearch()
	}
	return conf
}
*/

// serverOffset returns an offset that can be used to determine
// indices of servers in c.servers when making queries.
// When the rotate option is enabled, this offset increases.
// Otherwise it is always 0.
func (c *SPDNSConfig) serverOffset() uint32 {
	if c.Rotate {
		return atomic.AddUint32(&c.Soffset, 1) - 1 // return 0 to start
	}
	return 0
}

// nameList returns a list of names for sequential DNS queries.
func (conf *SPDNSConfig) nameList(name string) []string {
	if avoidDNS(name) {
		return nil
	}

	// Check name length (see isDomainName).
	l := len(name)
	rooted := l > 0 && name[l-1] == '.'
	if l > 254 || l == 254 && rooted {
		return nil
	}

	// If name is rooted (trailing dot), try only that name.
	if rooted {
		return []string{name}
	}

	hasNdots := count(name, '.') >= conf.Ndots
	name += "."
	l++

	// Build list of search choices.
	names := make([]string, 0, 1+len(conf.Search))
	// If name has enough dots, try unsuffixed first.
	if hasNdots {
		names = append(names, name)
	}
	// Try suffixes that are not too long (see isDomainName).
	for _, suffix := range conf.Search {
		if l+len(suffix) <= 254 {
			names = append(names, name+suffix)
		}
	}
	// Try unsuffixed, if not tried first above.
	if !hasNdots {
		names = append(names, name)
	}
	return names
}

func DnsDefaultSearch() []string {
	hn, err := getHostname()
	if err != nil {
		// best effort
		return nil
	}
	if i := strings.IndexByte(hn, '.'); i >= 0 && i < len(hn)-1 {
		return []string{ensureRooted(hn[i+1:])}
	}
	return nil
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func ensureRooted(s string) string {
	if len(s) > 0 && s[len(s)-1] == '.' {
		return s
	}
	return s + "."
}
