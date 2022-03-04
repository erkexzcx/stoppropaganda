package stoppropaganda

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

func fasthttpRequestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/status":
		fasthttpStatusResponseHandler(ctx)
	}
}

type StatusStruct struct {
	DNS      map[string]*DNSServerStatus `json:"DNS"`
	Websites map[string]*WebsiteStatus   `json:"Websites"`
}
type StatusService struct {
	AllStatus StatusStruct

	mux sync.Mutex
}

func fasthttpStatusResponseHandler(ctx *fasthttp.RequestCtx) {
	statusService := StatusService{
		AllStatus: StatusStruct{
			DNS:      make(map[string]*DNSServerStatus, len(dnsServers)),
			Websites: make(map[string]*WebsiteStatus, len(websites)),
		},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(dnsServers))
	wg.Add(len(websites))

	for endpoint, ds := range dnsServers {
		go func(endpoint string, ds *DNSServer) {
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
			ws.mux.Lock()
			tmpStatus := ws.Status
			unpauseTime := ws.unpauseTime
			ws.mux.Unlock()

			statusService.mux.Lock()
			statusService.AllStatus.Websites[endpoint] = &tmpStatus
			statusService.mux.Unlock()

			dosPausedFor := -time.Since(unpauseTime)
			if dosPausedFor > 0 {
				tmpStatus.Status += " for " + dosPausedFor.String()
			} else {
				tmpStatus.Status = "Running"
			}

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
