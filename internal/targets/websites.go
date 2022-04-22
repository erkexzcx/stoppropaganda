package targets

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

var TargetWebsites = map[string]struct{}{
	/* Other countries */

	"https://bukimevieningi.lt":                    {},
	"https://musutv.lt":                            {},
	"https://lt.baltnews.com":                      {},
	"http://sputniknews.lt":                        {},
	"https://lv.sputniknews.ru":                    {},
	"https://api.viada.lt":                         {},
	"https://www.sber.kz":                          {},
	"https://www.sberbank.kz":                      {},
	"https://online.sberbank.kz/CSAFront/index.do": {},

	/* Russia */

	// Propaganda
	"https://lenta.ru":                    {},
	"https://tass.ru":                     {},
	"https://tass.ru/userApi/getNewsFeed": {},
	"https://zakupki.gov.ru":              {},
	"https://www.rzd.ru":                  {},
	"https://rzdlog.ru":                   {},
	"https://www.interfax.ru":             {},
	"https://vz.ru":                       {},
	"https://sputniknews.ru":              {},
	"https://www.gazeta.ru":               {},
	"https://identity.kp.house/identity/api/2/auth/token.json": {},
	"https://api.riafan.ru":                      {},
	"https://pikabu.ru":                          {},
	"https://api.pikabu.ru":                      {},
	"https://i.kommersant.ru":                    {},
	"https://omk-job.ru":                         {},
	"https://www.rambler.ru":                     {},
	"https://www.rambler.ru/api/v4/personalized": {},
	// 	"https://id.rambler.ru/login-20/login": {}, // TODO fix small read buffer
	"https://mail.rambler.ru":                            {},
	"https://news.rambler.ru":                            {},
	"https://peroxide.rambler.ru/v1/comments/clusters/":  {},
	"https://rabota.rambler.ru":                          {},
	"https://rcm.rambler.ru":                             {},
	"https://mail.ru":                                    {},
	"https://www.astrobl.ru":                             {},
	"https://www2.astrobl.ru":                            {},
	"https://adm.astrobl.ru":                             {},
	"https://www.volgograd.ru":                           {},
	"https://www.admoblkaluga.ru":                        {},
	"https://mosreg.ru":                                  {},
	"https://easuz.mosreg.ru":                            {},
	"https://login.school.mosreg.ru":                     {},
	"https://uslugi.mosreg.ru":                           {},
	"https://edu-mosreg.ru":                              {},
	"https://orel-region.ru":                             {},
	"https://midural.ru":                                 {},
	"https://ulgov.ru":                                   {},
	"https://cheladmin.ru":                               {},
	"https://transport.cheladmin.ru":                     {},
	"https://stavregion.ru":                              {},
	"http://dostup.stavregion.ru/dsreda/login":           {},
	"http://lk.stavregion.ru":                            {},
	"https://torgi.stavregion.ru/application/":           {},
	"https://transport.stavregion.ru":                    {},
	"https://rg.ru":                                      {},
	"http://allpravda.info":                              {},
	"https://antifashist.com":                            {}, // https://antifashist.online
	"https://antimaydan.info":                            {},
	"https://argumenti.ru":                               {},
	"https://aurora.network":                             {},
	"https://balrad.ru":                                  {},
	"https://buzina.org":                                 {},
	"https://www.c-inform.info":                          {},
	"https://crimea-news.com":                            {},
	"https://crimeapress.info":                           {},
	"https://comitet.su":                                 {},
	"https://compromat.group":                            {},
	"https://compromat.group/engine/ajax/checkViews.php": {},
	"https://cont.ws":                                    {},
	"https://dan-news.info":                              {},
	"https://dni.ru":                                     {},
	"https://dnr24.com":                                  {},
	"https://dnr24.su":                                   {},
	"https://dnr-pravda.ru":                              {},
	"https://donbasstoday.ru":                            {},
	"https://dosie.su":                                   {}, // cf
	"https://eadaily.com/ru/dossier":                     {},
	"https://eadaily.com/ru/rss/index.xml":               {},
	"https://evening-crimea.com":                         {},
	"https://expert.ru":                                  {},
	"https://finobzor.ru":                                {},
	"https://forpostsevastopol.ru":                       {},
	"http://globalwarnews.ru":                            {},
	"https://gorlovka-news.su":                           {},
	"https://gtrklnr.com":                                {},
	"https://inforeactor.ru":                             {},
	"https://infox.sg":                                   {},
	"https://inlugansk.ru":                               {},
	"https://jpgazeta.ru":                                {},
	"https://kianews24.ru":                               {},
	"http://k-politika.ru":                               {},
	"https://kuban24.tv":                                 {},
	"https://life.ru":                                    {},
	"https://api.corr.life/public/regions/lookup":        {},
	"https://lug-info.com":                               {},
	"https://www.m24.ru":                                 {},
	"https://cross.m24.ru/covid/frontend/web/site/rus":   {},
	"https://mediarepost.ru":                             {},
	"https://www.metronews.ru":                           {},
	"https://www.readmetro.com":                          {},
	"https://www.mk-kalm.ru":                             {},
	"https://mkset.ru":                                   {},
	"http://www.moscow-post.su":                          {},
	"https://www.nakanune.ru":                            {},
	"https://nation-news.ru":                             {},
	"https://newizv.ru":                                  {},
	"https://newizv.ru/api/v1/matters":                   {},
	"http://news24today.info":                            {},
	"https://newsland.com":                               {},
	"https://nvo.ng.ru":                                  {},
	"https://nnews.nnov.ru":                              {},
	"https://newdaynews.ru":                              {},
	"https://www.newc.info":                              {},
	// 	"https://odnarodyna.org":                             {}, // TODO small read buffer
	"https://novorossiia.ru":              {}, // cf
	"http://novorossy.ru":                 {},
	"http://novosti333.ru":                {},
	"https://nt1941.su":                   {},
	"http://oddr.info":                    {},
	"https://pandoraopen.ru":              {},
	"https://plainnews.ru":                {},
	"https://pravdoryb.info":              {},
	"https://polit.info":                  {},
	"https://politcentr.ru":               {},
	"https://politexpert.net":             {},
	"https://politikus.ru":                {},
	"https://politpuzzle.ru":              {},
	"https://politros.com":                {},
	"https://rf-smi.ru":                   {},
	"https://www.ritmeurasia.org":         {},
	"https://rodina.news":                 {}, // cf
	"https://rossaprimavera.ru":           {},
	"https://rusnext.ru":                  {}, // cf
	"https://rusnext.ru/newslistapi.json": {}, // cf
	"http://ru-an.info":                   {},
	"https://rueconomics.ru":              {},
	"https://rusonline.org":               {},
	"http://ruspravda.info":               {},
	"https://slovodel.com":                {},
	"https://smi2.ru":                     {},
	"https://polls.smi2.ru/body/1/poll":   {},
	"https://sovross.ru":                  {},
	"http://stringer-news.com":            {},
	"https://svpressa.ru":                 {},
	"https://taurica.net":                 {},
	"https://tehnowar.ru":                 {},
	"https://time-news.net":               {}, // cf
	"https://trmzk.ru":                    {},
	"https://u-f.ru":                      {},
	"https://universe-tss.su":             {},
	"https://utro.ru":                     {},
	"https://www.vedomosti.ru":            {},
	"https://vesti-k.ru":                  {},
	"https://vm.ru":                       {},
	"https://voskhodinfo.su":              {},
	"https://webkamerton.ru":              {},
	"http://webnovosti.info":              {},
	"https://wpristav.ru":                 {}, // cf
	"https://x-true.info":                 {}, // cf
	"http://xvesti.ru":                    {},
	"https://zanamipravda.ru":             {},
	"https://zavtra.ru":                   {},

	// https://www.state.gov/russias-pillars-of-disinformation-and-propaganda-report/
	"http://www.geopolitika.ru":                 {},
	"https://katehon.com":                       {},
	"https://novorosinform.org":                 {},
	"https://www.pravda.ru":                     {},
	"https://www.pravda.ru/ajaxed/toolbartabs/": {},
	"https://russia-insider.com/ru":             {},
	"https://southfront.org":                    {},
	"https://www.strategic-culture.org":         {},
	"https://therussophile.org":                 {},
	"https://tsargrad.tv":                       {},
	"https://www.voltairenet.org":               {}, // cf

	// Business corporations
	"https://lukoil.ru":                                    {},
	"https://magnit.ru":                                    {},
	"https://www.evraz.com/ru/":                            {},
	"https://nlmk.com/ru/":                                 {},
	"https://www.severstal.com":                            {},
	"https://www.eurosib.ru":                               {},
	"https://www.wildberries.ru":                           {},
	"https://www.ozon.ru":                                  {},
	"https://www.ozon.ru/api/composer-api.bx/page/json/v2": {},
	"https://www.dns-shop.ru":                              {},
	"https://aliexpress.ru":                                {},
	"https://advego.com":                                   {},
	"https://freelance.ru":                                 {},
	"https://www.turbotext.ru":                             {},
	"https://sberfn.ru":                                    {},
	"https://www.vtbcapital-pr.ru":                         {},
	"https://region-am.ru":                                 {},
	"https://www.ingosinvest.ru":                           {},
	"https://goszakaz.ru":                                  {},
	"https://star-pro.ru":                                  {},
	"https://region.ru":                                    {},
	"https://gruzovozkin.pro":                              {},
	"https://ok.ru":                                        {},
	"http://5.61.23.11":                                    {},
	"http://217.20.155.13":                                 {},
	"http://217.20.147.1":                                  {},
	"http://www.yemelya.ru":                                {},
	"https://vcs.rostec.ru":                                {},
	"https://lk.rostec.ru/Account/LogIn":                   {},
	"http://ias.rostec.ru":                                 {},
	"https://smi.rostec.ru/user":                           {},
	"https://vks3.rostec.ru":                               {},
	"https://kontur.ru":                                    {},
	"https://help.kontur.ru/ke":                            {},

	// Banks
	"https://sber.ru": {},
	"https://api.developer.sber.ru/product/SberbankID": {},
	"https://promo.sber.ru/sberserenity":               {},
	"https://sberprime.sber.ru":                        {},
	"https://www.sberbank.com/ru":                      {},
	"https://www.sberbank.ru":                          {},
	"https://api.sberbank.ru/prod/tokens/v2":           {},
	"https://api.sberbank.ru/prod/tokens/v2/oauth":     {},
	"https://api.sberbank.ru/prod/tokens/v2/oidc":      {},
	"https://online.sberbank.ru/CSAFront/index.do":     {},
	"https://www.vtb.ru":                               {},
	"https://cl.vtb.ru":                                {},
	"https://ipoteka.vtb.ru/ipoteka/1":                 {},
	// 	"https://online.vtb.ru/login": {}, // TODO small read buffer
	"https://school.vtb.ru":                                          {},
	"https://siteapi.vtb.ru/api/sitepages/component-types":           {},
	"https://www.gazprombank.ru":                                     {},
	"https://api.gazprombank.ru":                                     {},
	"https://www.moex.com":                                           {},
	"https://iss.moex.com/iss/reference/":                            {},
	"https://messaging.moex.com/init":                                {},
	"https://passport.moex.com/login":                                {},
	"http://www.fsb.ru":                                              {},
	"https://scr.online.sberbank.ru/api/fl/idgib-w-3ds":              {},
	"https://3dsec.sberbank.ru/mportal3/auth/login":                  {},
	"https://acs1.sbrf.ru":                                           {},
	"https://acs2.sbrf.ru":                                           {},
	"https://acs3.sbrf.ru":                                           {},
	"https://acs4.sbrf.ru":                                           {},
	"https://acs5.sbrf.ru":                                           {},
	"https://acs6.sbrf.ru":                                           {},
	"https://acs7.sbrf.ru":                                           {},
	"https://acs8.sbrf.ru":                                           {},
	"https://my.bank-hlynov.ru/login/":                               {},
	"https://chbrr.crimea.com":                                       {},
	"https://online.vtb.ru":                                          {},
	"https://online.gpb.ru/login":                                    {},
	"https://alfabank.ru":                                            {},
	"https://alfabank.ru/api/v1/geco-ip/":                            {},
	"https://www.rshb.ru":                                            {},
	"https://online.rshb.ru/b1/ib6/wf2/retail/ib/loginretaildefault": {},
	"https://online.mkb.ru":                                          {},
	"https://www.tinkoff.ru":                                         {},
	"https://www.tinkoff.ru/api/common/v1/webuser":                   {},
	"https://id.tinkoff.ru/auth/step":                                {},
	"https://178.248.236.218":                                        {},
	"https://91.194.226.50":                                          {},
	"https://auth.tcsbank.ru":                                        {},
	"https://91.194.226.32":                                          {},

	//The state
	"https://www.gosuslugi.ru":                    {},
	"https://www.gosuslugi.ru/api/mainpage/v4":    {},
	"https://www.mos.ru/uslugi":                   {},
	"https://api.mos.ru":                          {},
	"http://kremlin.ru":                           {},
	"http://government.ru":                        {},
	"https://mil.ru":                              {},
	"https://www.nalog.gov.ru":                    {},
	"https://customs.gov.ru":                      {},
	"https://pfr.gov.ru":                          {},
	"https://gosuslugi41.ru":                      {},
	"https://uslugi27.ru":                         {},
	"https://gosuslugi29.ru":                      {},
	"https://www.pochta.ru":                       {},
	"https://otpravka.pochta.ru":                  {},
	"https://passport.pochta.ru/oauth2/authorize": {},
	"https://tariff.pochta.ru":                    {},
	"https://zakaznoe.pochta.ru":                  {},
	"http://crimea-post.ru":                       {},
	"https://ca.vks.rosguard.gov.ru":              {},
	"https://sozd.duma.gov.ru":                    {},
	"https://vsednr.ru":                           {},
	"https://wagnera.ru":                          {},

	// Embassies
	// Do not duplicate the IPs: https://github.com/erkexzcx/stoppropaganda/pull/110#issuecomment-1059960305
	"https://montreal.mid.ru": {},
	"https://belarus.mid.ru":  {},

	// Others
	"https://mail.rkn.gov.ru":    {},
	"https://cloud.rkn.gov.ru":   {},
	"https://proverki-gov.ru":    {},
	"https://www.glonass-iac.ru": {},
	"https://dme.ru":             {},
	"https://market.dme.ru":      {},
	"https://www.tickets.dme.ru": {},

	// Payments
	"http://185.170.2.7":                  {},
	"http://185.170.3.7":                  {},
	"http://185.170.2.9":                  {},
	"https://ds1.mirconnect.ru":           {},
	"https://ds2.mirconnect.ru":           {},
	"https://uat-ds1.mirconnect.ru":       {},
	"https://uat-ds2.mirconnect.ru":       {},
	"https://sberpay.mirconnect.ru":       {},
	"https://wats.mirconnect.ru":          {},
	"https://ds.vendorcert.mirconnect.ru": {},
	"https://privetmir.ru":                {},
	"https://sbp.nspk.ru":                 {},
	"https://sip.nspk.ru":                 {},
	"https://koronapay.com":               {},
	"https://api.koronapay.com":           {},
	"https://etpgpb.ru":                   {},
	"https://api.etpgpb.ru":               {},
	"https://etp.gpb.ru":                  {},
	"https://passport.etpgpb.ru":          {},
	"https://gos.etpgpb.ru":               {},
	"https://inter-rao.etpgpb.ru":         {},
	"https://rosnedra.etpgpb.ru":          {},
	"https://gpn.etpgpb.ru":               {},
	"https://geh.etpgpb.ru":               {},
	"https://trade.etpgpb.ru":             {},
	"https://gb.etpgpb.ru":                {},
	"https://bki-okb.ru":                  {},

	// Exchanges connected to russian banks
	"https://bitokk.biz":              {},
	"https://www.netex24.net":         {}, // requires cookie through redirect to fully load
	"https://flashobmen.com":          {},
	"https://yoomoney.ru":             {},
	"https://yookassa.ru":             {},
	"https://telemetry.koronapay.com": {},
	"https://olympus.koronapay.com":   {},

	// Electronic signature services, certificate authorities, www domain names
	"http://www.nucrf.ru":                 {},
	"http://www.belinfonalog.ru":          {},
	"https://www.roseltorg.ru":            {},
	"https://178fz.roseltorg.ru":          {},
	"https://agro.roseltorg.ru":           {},
	"https://atom2.roseltorg.ru":          {},
	"https://bg.roseltorg.ru/auth/email":  {},
	"https://com.roseltorg.ru":            {},
	"https://docs.roseltorg.ru":           {},
	"https://etp.roseltorg.ru":            {},
	"https://fkr.roseltorg.ru":            {},
	"https://fkr2.roseltorg.ru":           {},
	"https://kim-atom.roseltorg.ru":       {},
	"https://kim.roseltorg.ru":            {},
	"https://kim-irao.roseltorg.ru":       {},
	"https://lk.roseltorg.ru/auth":        {},
	"https://orders.roseltorg.ru":         {},
	"https://rosgeo.roseltorg.ru":         {},
	"https://rosseti.roseltorg.ru":        {},
	"https://rt.roseltorg.ru":             {},
	"https://rushydro.roseltorg.ru":       {},
	"https://msp.roseltorg.ru":            {},
	"https://vtb.roseltorg.ru":            {},
	"https://rosinvoice.ru/auth":          {},
	"http://www.nwudc.ru":                 {},
	"https://kk.bank":                     {},
	"https://kb.kubankredit.ru/login":     {},
	"https://structure.mil.ru":            {},
	"http://www.ucpir.ru":                 {},
	"http://www.e-portal.ru":              {},
	"https://api.e-portal.ru":             {},
	"https://parus-s.ru":                  {},
	"https://www.icentr.ru":               {},
	"https://etp.kartoteka.ru/index.html": {},
	"http://rsbis.ru":                     {},
	"https://www.24ecp.ru":                {},
	"https://kraskript.com":               {},
	"http://ca.ntssoft.ru":                {},
	"http://www.y-center.ru":              {},
	"http://www.rcarus.ru":                {},
	"https://squaretrade.ru":              {},
	"http://www.otchet-online.ru":         {},
	"http://udcs.ru":                      {},
	"https://www.cit-ufa.ru":              {},
	"https://api.cit-ufa.ru":              {},
	"http://elkursk.ru":                   {},
	"https://icvibor.ru":                  {},
	"https://1c.icvibor.ru":               {},
	"http://ucestp.ru":                    {},
	"https://api.ucestp.ru":               {},
	"http://mcspro.ru":                    {},
	"http://www.infotrust.ru":             {},
	"http://epnow.ru":                     {},
	"http://cfmc.ru":                      {},
	"https://esia.gosuslugi.ru":           {},
	"https://iecp.ru/ep/ep-verification":  {},
	"https://e-trust.gosuslugi.ru":        {},
	"https://gu.spb.ru":                   {},
	"https://iecp.ru/ep/uc-list":          {},

	// Oil and gas trading companies
	"https://www.tektorg.ru":                                      {},
	"https://44.tektorg.ru/authentication/login":                  {},
	"https://fin.tektorg.ru/bg":                                   {}, // static
	"https://fin.tektorg.ru/api/v2/reference/service_types":       {}, // api pass
	"https://zakupki.tektorg.ru":                                  {},
	"https://rosneft.com":                                         {},
	"https://lukoil.com":                                          {},
	"https://aaa-srt-cs.lukoil.com/logon/LogonPoint/tmindex.html": {},
	"https://ar.lukoil.com/login":                                 {},
	"https://gazprom.ru":                                          {},
	"https://b2b.sibur.ru/pages_new_en/faces/index.jspx":          {},
	"https://onlinecontract.ru":                                   {},
	"https://www.eurochemgroup.com":                               {},
	"https://tatneft.ru":                                          {},

	// Food delivery services
	"https://www.delivery-club.ru":                    {},
	"https://api.delivery-club.ru/api1.2/experiments": {},
	"https://m-food.ru":                               {},
	"https://sbermarket.ru":                           {},
	"https://chibbis.ru":                              {},
	"https://samokat.ru":                              {},
	"https://localkitchen.ru":                         {},

	// Cinemas
	"https://cinemastar.ru":     {},
	"https://kinoteatr.ru":      {},
	"https://karofilm.ru":       {},
	"https://romanov-cinema.ru": {},
	"https://www.mirage.ru":     {},

	// Belgorod
	"https://bel.ru": {},
	"https://bel.ru/api/v1/platform/main_page":                 {},
	"https://belgorod.igooods.ru":                              {},
	"https://belgorod.igooods.ru/api/v8/cities/identify_by_ip": {},
	"https://belgorod.farfor.ru":                               {},
	"https://belsbyt.ru":                                       {},
	"https://lk.belsbyt.ru/fiz/login":                          {},
	"https://www.go31.ru":                                      {},
	"https://www.go31.ru/api3/auth/check-login":                {},
	"https://perevozki31.ru":                                   {},

	// Others
	"http://217.12.104.100":  {},
	"https://toxic-repos.ru": {},

	// Various websites by ip
	"https://91.213.144.193": {},
	"https://91.213.144.237": {},
	"https://78.47.115.99":   {},
	"https://81.19.72.39":    {},
	"https://81.19.72.3":     {},
	"https://194.190.37.226": {},
	"https://194.190.37.228": {},
	"https://194.190.37.234": {},
	"https://213.59.254.8":   {},
	"https://185.194.34.142": {},
	"https://213.59.197.65":  {},
	"https://185.71.67.237":  {},
	"https://82.151.111.187": {},
	"https://91.239.5.75":    {},
	"https://185.58.223.206": {},
	"https://178.248.238.15": {},
	"https://92.53.98.191":   {},

	/* BELARUS */

	// by gov
	"https://mininform.gov.by":       {},
	"https://rec.gov.by/ru":          {},
	"https://www.mil.by":             {},
	"https://www.government.by":      {},
	"https://president.gov.by/ru":    {},
	"https://www.mvd.gov.by/ru":      {},
	"https://www.kgb.by/ru":          {},
	"https://www.prokuratura.gov.by": {},
	"https://mfa.gov.by":             {},
	"https://russia.mfa.gov.by/ru":   {},

	// by banks
	"https://www.nbrb.by":                 {},
	"https://belarusbank.by":              {},
	"https://brrb.by":                     {},
	"https://www.belapb.by":               {},
	"https://bankdabrabyt.by":             {},
	"https://fx.bankdabrabyt.by":          {},
	"https://online.bankdabrabyt.by":      {},
	"https://belinvestbank.by/individual": {},
	"https://api.belinvestbank.by":        {},
	"https://belpost.by":                  {},
	"https://api.belpost.by/api/v1/pages": {},

	// by business
	"https://bgp.by/ru":           {},
	"https://www.belneftekhim.by": {},
	"http://www.bellegprom.by":    {},
	"https://www.energo.by":       {},
	"http://belres.by/ru":         {},
	"https://www.rw.by":           {},

	// by media
	"https://www.belta.by":        {},
	"https://sputnik.by":          {},
	"https://www.tvr.by":          {},
	"https://www.sb.by":           {},
	"https://www.belarus.by":      {},
	"https://belarus24.by":        {},
	"https://ont.by":              {},
	"https://www.024.by":          {},
	"https://www.belnovosti.by":   {},
	"https://mogilevnews.by":      {},
	"https://www.slonves.by":      {},
	"http://www.ctv.by":           {},
	"https://radiobelarus.by":     {},
	"https://radiusfm.by":         {},
	"https://radiomir.by":         {},
	"https://api.radiomir.by":     {},
	"https://radiostalica.by":     {},
	"https://radiobrestfm.by":     {},
	"https://api.radiobrestfm.by": {},
	"https://www.tvrmogilev.by":   {},
	"https://zarya.by":            {},
	"https://grodnonews.by":       {},

	/* Syria */
	"https://syrianfinance.gov.sy": {},
	"http://185.216.132.201":       {},

	/* DDOS mitigation */
	"https://ddos-guard.net/ru": {},
	"https://stormwall.pro":     {},
	"https://qrator.net/ru":     {},
	"https://solidwall.ru":      {},

	// Still operating in Russia
	// https://www.yalerussianbusinessretreat.com/
	"http://api.auchan.ru":                                          {},
	"http://api-partners.auchan.ru":                                 {},
	"https://transport.auchan.ru/core/framework/login.cfm":          {},
	"https://burgerkingrus.ru":                                      {},
	"https://burgerkingrus.ru/api-web-front/api/v3/restaurant/list": {},
	"https://mic.burgerking.ru/mifs/user/login.jsp":                 {},
	"https://sd.burgerking.ru/HEAT/SaaSExternalSessionRenew.aspx":   {},
	"https://www.decathlon.ru":                                      {},
	"https://blog.decathlon.ru":                                     {},
	"https://delonghi.ru":                                           {},
	"https://ecco.ru":                                               {},
	"https://www.jd.ru":                                             {},
	"https://www.jd.ru/leproxy/api/carts/cart/owner-id":             {},
	"https://mega.ru":                                               {}, // Ingka
	"https://modshairrussia.ru":                                     {},
	"https://mymodshair.ru":                                         {},
	"https://natr.ru":                                               {},
	"https://papajohns.ru/moscow":                                   {},
	"https://api.papajohns.ru/slider/list":                          {},
	"https://subway.ru":                                             {},
	"https://www.yves-rocher.ru":                                    {},

	// https://t.me/itarmyofukraine2022/230
	"http://www.ved.gov.ru":      {},
	"https://www.mid.ru":         {},
	"https://www.economy.gov.ru": {},
	"https://minobrnauki.gov.ru": {},

	// https://t.me/itarmyofukraine2022/215
	"https://www.chechnya.online": {},
	"https://vaynahavia.com":      {},
	"https://www.minfinchr.ru":    {},

	"https://rostelecom.ru":     {},
	"https://www.company.rt.ru": {},

	// https://t.me/itarmyofukraine2022/243
	"https://fss.gov.ru":              {},
	"https://autodiscover.fss.ru/owa": {},
	"https://portal.fss.ru":           {},
	"https://data.fss.ru/open":        {},
	"http://docs.fss.ru":              {},
	"https://hosting.pfrf.ru":         {},
	"https://school.pfr.gov.ru":       {},
	"https://es-pfrf.ru":              {},

	// https://t.me/itarmyofukraine2022/246
	"https://rosleshoz.gov.ru": {},

	// https://t.me/itarmyofukraine2022/248
	"https://digital.ac.gov.ru":       {},
	"https://lk.digital.gov.ru":       {},
	"https://platform.digital.gov.ru": {},
	"https://reestr.digital.gov.ru":   {},
	"https://sc.digital.gov.ru/home":  {},
	"https://minenergo.gov.ru":        {},
	"https://minjust.gov.ru":          {},
	"https://www.minsport.gov.ru":     {},
	"https://minstroyrf.gov.ru":       {},
	"https://mintrud.gov.ru":          {},

	// https://t.me/itarmyofukraine2022/251
	"https://payanyway.ru/info/w/ru/public/welcome.htm": {},

	// https://t.me/itarmyofukraine2022/253
	"https://www.superjob.ru": {}, // requires cookie through redirect to fully load
	"https://rabota.vk.com":   {},

	// https://t.me/itarmyofukraine2022/269
	"https://cdek.by/ru/":         {},
	"https://www.cdek.ru/graphql": {},

	// https://t.me/itarmyofukraine2022/273
	"https://www.doski.ru":                           {},
	"https://www.farpost.ru":                         {},
	"https://unibo.ru":                               {},
	"https://api.youla.io/api/v1/search/suggestions": {},

	// https://t.me/itarmyofukraine2022/275
	"https://www.crpt.ru":                  {},
	"https://ismp.crpt.ru":                 {},
	"https://kb.crpt.ru":                   {},
	"https://markirovka.crpt.ru/login-kep": {},
	"https://api.mdlp.crpt.ru":             {},
	"https://api.sb.mdlp.crpt.ru":          {},
	"https://nrz.api.sb.mdlp.crpt.ru":      {},
	"https://milk.crpt.ru":                 {},
	"https://suzgrid.crpt.ru":              {},

	// https://t.me/itarmyofukraine2022/287
	"https://www.amediateka.ru": {},
	"https://www.okko.tv":       {},
	"https://www.ontvtime.ru":   {},
	"https://wink.ru":           {},

	// https://t.me/itarmyofukraine2022/289
	"https://stream.1tv.ru/api/playlist/1tvch-v1_as_array.json": {},

	// https://t.me/itarmyofukraine2022/290
	"https://api.kontur.ru/csi/support/v1/public/integrations/zakupki": {},
	"https://auth.kontur.ru":                       {},
	"https://ca.kontur.ru":                         {},
	"https://developer.kontur.ru":                  {},
	"https://elbank.kontur.ru/AccessControl/Login": {},
	"https://focus.kontur.ru":                      {},
	"https://focus.kontur.ru/api/lists/all":        {},
	"https://install.kontur.ru/kekep":              {},
	"https://pf.kontur.ru":                         {},
	"https://zakupki.kontur.ru":                    {},
	"https://www.b-kontur.ru":                      {},
	"https://www.kontur-extern.ru":                 {},

	// https://t.me/itarmyofukraine2022/302
	"https://rov.aero": {},
}
