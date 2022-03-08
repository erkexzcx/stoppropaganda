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
	"217.175.155.12:53":  {},
	"217.175.140.71:53":  {},
}

// We need reliable IP addresses
// just like we would've been in Russia/Belarus
var ReferenceDNSServersForHTTP = []string{
	// https://dns.yandex.com/
	"77.88.8.8:53",
	"77.88.8.1:53",
}
