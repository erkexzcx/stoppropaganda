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
	"217.12.96.15:53":    {}, // ns1.alfadirect.net
	"217.12.97.15:53":    {}, // ns2.alfadirect.net
	"217.12.102.15:53":   {}, // ns1.beta-bank.ru
	"217.12.102.16:53":   {}, // ns2.beta-bank.ru
	"212.109.207.140:53": {}, // ns1.bank-hlynov.ru
	"168.63.61.187:53":   {}, // ns2.bank-hlynov.ru
	"46.22.48.162:53":    {}, // dns.ecco.ru
	"193.148.44.179:53":  {}, // ns.fss.ru
	"193.148.44.180:53":  {}, // ns2.fss.ru
	"95.173.128.77:53":   {}, // ns.gov.ru
	"95.173.130.4:53":    {}, // ns1.duma.gov.ru
	"95.173.130.5:53":    {}, // ns2.duma.gov.ru
	"94.124.200.34:53":   {}, // ns.hh.ru
	"185.12.92.62:53":    {}, // ns1.journal-neo.org
	"74.119.194.66:53":   {}, // ns2.journal-neo.org
	"185.170.2.237:53":   {}, // ns01.mirconnect.ru MIR payment system
	"185.170.3.237:53":   {}, // ns02.mirconnect.ru MIR payment system
	"195.209.134.8:53":   {}, // ns03.mirconnect.ru MIR payment system
	"195.209.181.8:53":   {}, // ns04.mirconnect.ru MIR payment system
	"195.161.52.77:53":   {}, // ns1.pfrf.ru
	"195.161.53.77:53":   {}, // ns2.pfrf.ru
	"79.142.16.28:53":    {}, // ns2.qiwi.com
	"79.142.17.8:53":     {}, // ns3.qiwi.com
	"79.142.17.28:53":    {}, // ns4.qiwi.com
	"79.142.16.8:53":     {}, // ns5.qiwi.com
	"81.19.73.8:53":      {}, // ns2.rambler.ru
	"81.19.83.8:53":      {}, // ns3.rambler.ru
	"81.19.73.9:53":      {}, // ns4.rambler.ru
	"81.19.83.9:53":      {}, // ns5.rambler.ru
	"195.93.247.59:53":   {}, // ns7.sputniknews.ru
	"195.93.247.60:53":   {}, // ns8.sputniknews.ru
	"178.248.236.20:53":  {}, // ns9.sputniknews.ru
	"178.248.233.32:53":  {}, // ns10.sputniknews.ru
	"77.238.100.195:53":  {}, // ns1.tsargrad.tv
	"188.64.160.163:53":  {}, // ns2.tsargrad.tv
}

// We need reliable IP addresses
// just like we would've been in Russia/Belarus
var ReferenceDNSServersForHTTP = []string{
	// https://dns.yandex.com/
	"77.88.8.8:53",
	"77.88.8.1:53",
	"77.88.8.2:53",
	"77.88.8.3:53",
	"77.88.8.7:53",
	"77.88.8.8:53",
	"77.88.8.88:53",
}
