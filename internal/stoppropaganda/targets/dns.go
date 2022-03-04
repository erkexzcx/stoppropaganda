package targets

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

var TargetDNSServers = map[string]struct{}{
	"194.54.14.186:53":  {},
	"194.54.14.187:53":  {},
	"194.67.7.1:53":     {},
	"194.67.2.109:53":   {},
	"84.252.147.118:53": {},
	"84.252.147.119:53": {},
	"95.173.148.51:53":  {},
	"95.173.148.50:53":  {},
}

// We need to get reliable IP address
// just like we would've been in Russia/Belarus
var ReferenceDNSServersForHTTP = []string{
	// https://dns.yandex.com/
	"77.88.8.8:53",
	"77.88.8.1:53",
}
