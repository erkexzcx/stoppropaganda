package stoppropaganda

import (
	"strconv"
	"strings"
	"time"

	"github.com/erkexzcx/stoppropaganda/internal/stoppropaganda/sockshttp"
)

// func CustomDial(addr string) (conn net.Conn, err error) {
// 	dialer, err := MakeDialerThrough(proxyChain)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

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
