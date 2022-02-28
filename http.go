package stoppropaganda

import (
	"encoding/json"
	"net/http"
	"sync"
)

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
