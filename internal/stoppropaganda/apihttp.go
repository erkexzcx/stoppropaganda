package stoppropaganda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/erkexzcx/stoppropaganda/internal/customresolver"
	"github.com/valyala/fasthttp"
)

func fasthttpRequestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/status":
		fasthttpStatusResponseHandler(ctx)
	case "/dnscache":
		fasthttpDnsCacheResponseHandler(ctx)
	case "/downloaded":
		fasthttpDownloadedResponseHandler(ctx)
	}
}

type StatusStruct struct {
	DNS      map[string]*DNSTargetStatus `json:"DNS"`
	Websites map[string]*WebsiteStatus   `json:"Websites"`
}
type StatusService struct {
	AllStatus StatusStruct

	mux sync.Mutex
}

func fasthttpStatusResponseHandler(ctx *fasthttp.RequestCtx) {
	statusService := StatusService{
		AllStatus: StatusStruct{
			DNS:      make(map[string]*DNSTargetStatus, len(dnsTargets)),
			Websites: make(map[string]*WebsiteStatus, len(websites)),
		},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(dnsTargets))
	wg.Add(len(websites))

	for endpoint, ds := range dnsTargets {
		go func(endpoint string, ds *DNSTarget) {
			ds.mux.Lock()
			dnsStatus := ds.Status
			ds.mux.Unlock()

			statusService.mux.Lock()
			statusService.AllStatus.DNS[endpoint] = &dnsStatus
			statusService.mux.Unlock()

			wg.Done()
		}(endpoint, ds)
	}

	for endpoint, ws := range websites {
		go func(endpoint string, ws *Website) {
			ws.statusMux.Lock()
			tmpStatus := ws.status
			ws.statusMux.Unlock()

			statusService.mux.Lock()
			statusService.AllStatus.Websites[endpoint] = &tmpStatus
			statusService.mux.Unlock()

			wg.Done()
		}(endpoint, ws)
	}

	wg.Wait()

	statusService.mux.Lock()
	content, err := json.MarshalIndent(statusService.AllStatus, "", "    ")
	statusService.mux.Unlock()
	if err != nil {
		ctx.SetStatusCode(500)
		ctx.WriteString("failed to marshal data")
		return
	}
	ctx.Write(content)
}

func fasthttpDnsCacheResponseHandler(ctx *fasthttp.RequestCtx) {

	cache := customresolver.DnsCache

	dnsCacheItems := cache.Items()
	content, err := json.MarshalIndent(dnsCacheItems, "", "    ")
	if err != nil {
		ctx.SetStatusCode(500)
		ctx.WriteString("failed to marshal data")
		return
	}
	ctx.Write(content)
}

type DownloadedStat struct {
	Endpoint   string
	Downloaded uint64
}

func (ds *DownloadedStat) FormatMegabytes() string {
	return fmt.Sprintf("%9.3f", float64(ds.Downloaded)/(1024.0*1024.0))
}

type DownloadedStats []*DownloadedStat

func (s DownloadedStats) Len() int      { return len(s) }
func (s DownloadedStats) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s DownloadedStats) Less(i, j int) bool {
	a := s[i].Downloaded
	b := s[j].Downloaded
	if a == b {
		return s[i].Endpoint < s[j].Endpoint
	}
	return a > b
}

func fasthttpDownloadedResponseHandler(ctx *fasthttp.RequestCtx) {
	stats := make([]*DownloadedStat, 0, len(websites))
	for endpoint, website := range websites {
		website.statusMux.Lock()
		downloaded := website.status.Downloaded
		website.statusMux.Unlock()
		stat := &DownloadedStat{
			Endpoint:   endpoint,
			Downloaded: downloaded,
		}
		stats = append(stats, stat)
	}

	sort.Sort(DownloadedStats(stats))

	buf := new(bytes.Buffer)

	if ctx.URI().QueryArgs().Has("raw") {
		for _, stat := range stats {
			fmt.Fprintf(buf, "%d\t%s\n", stat.Downloaded, stat.Endpoint)
		}
	} else {
		for _, stat := range stats {
			fmt.Fprintf(buf, "%s MB\t%s\n", stat.FormatMegabytes(), stat.Endpoint)
		}
	}

	ctx.Write(buf.Bytes())
}
