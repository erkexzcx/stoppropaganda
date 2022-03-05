package stoppropaganda

import (
	"bytes"
	"strconv"
	"strings"
	"time"

	"github.com/erkexzcx/stoppropaganda/internal/sockshttp"
)

const (
	ProxyMethodSocks5 = byte(iota)
	ProxyMethodSocks4
	ProxyMethodHttp
	ProxyMethodDirect
)

func MakeDialerThrough(parentDialer sockshttp.Dialer, proxyChain ProxyChain, proxyTimeout time.Duration) (dialer sockshttp.Dialer) {
	dialer = parentDialer
	for _, proxy := range proxyChain {
		proxyaddr := proxy.Addr
		method := proxy.Method
		if method == ProxyMethodDirect {
			// direct
		} else if method == ProxyMethodHttp {
			httpd, _ := sockshttp.HTTP("tcp", proxyaddr, dialer)
			dialer = httpd
			httpd.(*sockshttp.Http).Timeout = proxyTimeout
		} else if method == ProxyMethodSocks5 {
			socks5d, _ := sockshttp.SOCKS5("tcp", proxyaddr, nil, dialer)
			dialer = socks5d
			socks5d.(*sockshttp.Socks5).Timeout = proxyTimeout
		} else if method == ProxyMethodSocks4 {
			socks4d, _ := sockshttp.SOCKS4("tcp", proxyaddr, nil, dialer)
			dialer = socks4d
			socks4d.(*sockshttp.Socks4).Timeout = proxyTimeout

		}
	}

	return
}

func MethodName2ID(name string) byte {
	if name == "socks" || name == "socks5" || name == "5" {
		return ProxyMethodSocks5
	}
	if name == "http" {
		return ProxyMethodHttp
	}
	if name == "socks4" || name == "4" {
		return ProxyMethodSocks4
	}
	return ProxyMethodDirect
}
func MethodID2Name(id byte) string {
	if id == ProxyMethodSocks5 {
		return "socks5"
	}
	if id == ProxyMethodHttp {
		return "http"
	}
	if id == ProxyMethodSocks4 {
		return "socks4"
	}
	if id == ProxyMethodDirect {
		return "direct"
	}
	return "unknown " + strconv.Itoa(int(id))
}

type Proxy struct {
	Addr   string
	Method byte
}

func (p Proxy) String() string {
	return p.Addr
}

type ProxyChain []Proxy

func (pc ProxyChain) String() string {
	b := new(strings.Builder)
	for i, p := range []Proxy(pc) {
		if i != 0 {
			b.WriteByte(',')
		}
		b.WriteString(p.Addr)
	}
	return b.String()
}
func (pc ProxyChain) Last() Proxy {
	return pc[len(pc)-1]
}

func MaybeExtract(proxyDesc []byte, prefix []byte, methodIn byte) (ok bool, proxyAddr []byte, method byte) {
	ok = bytes.HasPrefix(proxyDesc, prefix)
	if ok {
		method = methodIn
		proxyAddr = proxyDesc[len(prefix):]
	}
	return
}
func ExtractProxyMethod(proxyDesc []byte) (proxyAddr []byte, method byte) {
	method = ProxyMethodSocks5

	if bytes.Equal(proxyDesc, []byte("direct")) {
		return proxyDesc, ProxyMethodDirect
	}

	var ok bool
	if ok, proxyAddr, method = MaybeExtract(proxyDesc, []byte("4:"), ProxyMethodSocks4); ok {
		return
	}
	if ok, proxyAddr, method = MaybeExtract(proxyDesc, []byte("socks4:"), ProxyMethodSocks4); ok {
		return
	}
	if ok, proxyAddr, method = MaybeExtract(proxyDesc, []byte("h:"), ProxyMethodHttp); ok {
		return
	}
	if ok, proxyAddr, method = MaybeExtract(proxyDesc, []byte("http:"), ProxyMethodHttp); ok {
		return
	}
	if ok, proxyAddr, method = MaybeExtract(proxyDesc, []byte("socks5:"), ProxyMethodSocks5); ok {
		return
	}
	if ok, proxyAddr, method = MaybeExtract(proxyDesc, []byte("5:"), ProxyMethodSocks5); ok {
		return
	}
	proxyAddr = proxyDesc

	return
}
func ParseProxy(proxyDesc string) (proxy Proxy) {
	proxyaddrBytes, method := ExtractProxyMethod([]byte(proxyDesc))
	proxyaddr := string(proxyaddrBytes)
	proxy = Proxy{Addr: proxyaddr, Method: method}
	return
}

func ParseProxyChain(proxyChain string) (chain ProxyChain) {
	// "socks4://1.1.1.1:1,h://2.2.2.2:2"
	proxyChain = strings.Replace(proxyChain, "//", "", -1)
	// "socks4:1.1.1.1:1,h:2.2.2.2:2"

	proxyAddrs := strings.Split(proxyChain, ",")
	// "socks4:1.1.1.1:1"
	// "h:2.2.2.2:2"

	chain = make([]Proxy, 0, len(proxyAddrs))
	for _, proxyDesc := range proxyAddrs {
		if proxyDesc != "direct" {
			proxy := ParseProxy(proxyDesc)
			chain = append(chain, proxy)
		}
	}
	return
}
