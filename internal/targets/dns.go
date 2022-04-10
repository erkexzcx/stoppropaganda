package targets

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

var TargetDNSServers = map[string]struct{}{
	"193.232.128.6:53":   {}, // ru root server
	"194.85.252.62:53":   {}, // ru root server
	"194.190.124.17:53":  {}, // ru root server
	"193.232.142.17:53":  {}, // ru root server
	"193.232.156.17:53":  {}, // ru root server
	"194.54.14.186:53":   {}, // Sberbank of Russia
	"194.54.14.187:53":   {}, // Sberbank of Russia
	"194.67.7.1:53":      {}, // Metal and Mineral (except Petroleum) Merchant Wholesalers Industry
	"194.67.2.109:53":    {}, // Teleross (golden telecom)
	"84.252.147.118:53":  {}, // 84.* => Sherbank public services network
	"84.252.147.119:53":  {},
	"95.173.148.51:53":   {}, // 95.* => Federal Guard Service of the Russian Federation
	"95.173.148.50:53":   {},
	"217.175.155.100:53": {},
	"92.53.97.198:53":    {},
	"195.10.198.37:53":   {},
	"185.170.2.237:53":   {}, // ns01.mirconnect.ru MIR payment system
	"185.170.3.237:53":   {}, // ns02.mirconnect.ru MIR payment system
	"195.209.134.8:53":   {}, // ns03.mirconnect.ru MIR payment system
	"195.209.181.8:53":   {}, // ns04.mirconnect.ru MIR payment system
	"193.148.44.179:53":  {}, // ns.fss.ru
	"193.148.44.180:53":  {}, // ns2.fss.ru
	"95.173.128.77:53":   {}, // ns.gov.ru
	"195.161.52.77:53":   {}, // ns1.pfrf.ru
	"195.161.53.77:53":   {}, // ns2.pfrf.ru
	"94.124.200.34:53":   {}, // ns.hh.ru
	"91.230.251.64:53":   {}, // ns01.crpt.ru
	"91.208.238.64:53":   {}, // ns03.crpt.ru
}

// We need reliable IP addresses
// just like we would've been in Russia/Belarus
var ReferenceDNSServersForHTTP = []string{
	// https://dns.yandex.com/
	"77.88.8.8:53",
	"77.88.8.1:53",
}
