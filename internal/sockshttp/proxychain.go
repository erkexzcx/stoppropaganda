package sockshttp

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

const (
	ProxyMethodSocks5 = byte(iota)
	ProxyMethodSocks4
	ProxyMethodHttp
	ProxyMethodDirect
	ProxyMethodUnknown
)

func MakeDialerThrough(parentDialer Dialer, proxyChain ProxyChain, proxyTimeout time.Duration) (dialer Dialer) {
	dialer = parentDialer
	for _, proxy := range proxyChain {
		proxyaddr := proxy.Addr
		method := proxy.Method
		auth := proxy.Auth
		if method == ProxyMethodDirect {
			// direct
		} else if method == ProxyMethodHttp {
			httpd, _ := HTTP("tcp", proxyaddr, &auth, dialer)
			httpd.Timeout = proxyTimeout
			dialer = httpd
		} else if method == ProxyMethodSocks5 {
			socks5d, _ := SOCKS5("tcp", proxyaddr, &auth, dialer)
			socks5d.Timeout = proxyTimeout
			dialer = socks5d
		} else if method == ProxyMethodSocks4 {
			socks4d, _ := SOCKS4("tcp", proxyaddr, &auth, dialer)
			socks4d.Timeout = proxyTimeout
			dialer = socks4d
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
	return ProxyMethodUnknown
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
	Auth   Auth
}

func (p Proxy) String() string {
	return MethodID2Name(p.Method) + ":" + p.Addr
}

type ProxyChain []Proxy

func (pc ProxyChain) String() string {
	b := new(strings.Builder)
	for i, p := range []Proxy(pc) {
		if i != 0 {
			b.WriteByte(',')
		}
		b.WriteString(p.String())
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

	index := bytes.IndexByte(proxyDesc, byte(':'))
	if index != -1 {
		method = MethodName2ID(string(proxyDesc[:index]))
		if method != ProxyMethodUnknown {
			proxyAddr = proxyDesc[index+1:]
		} else {
			// Proxy didn't have prefix like "socks4://"
			// so let's default it to socks5
			method = ProxyMethodSocks5
			proxyAddr = proxyDesc
		}
	} else {
		proxyAddr = proxyDesc
	}
	return
}
func ParseProxy(proxyDesc string) (proxy Proxy) {

	proxyaddrBytes, method := ExtractProxyMethod([]byte(proxyDesc))
	proxyaddr := string(proxyaddrBytes)
	auth, remaining := parseAuthority(proxyaddr)
	proxy = Proxy{Addr: remaining, Method: method, Auth: auth}
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

func parseAuthority(authority string) (auth Auth, remaining string) {
	i := strings.LastIndex(authority, "@")
	if i < 0 {
		remaining = authority
		return
	} else {
		remaining = authority[i+1:]
	}

	userinfo := authority[:i]
	if !validUserinfo(userinfo) {
		return
	}
	if !strings.Contains(userinfo, ":") {
		auth = Auth{
			User:     userinfo,
			Password: "",
		}
	} else {
		splitted := strings.SplitN(userinfo, ":", 2)
		username := splitted[0]
		password := ""
		if len(splitted) > 1 {
			password = splitted[1]
		}
		auth = Auth{
			User:     username,
			Password: password,
		}
	}
	return
}

// validUserinfo reports whether s is a valid userinfo string per RFC 3986
// Section 3.2.1:
//     userinfo    = *( unreserved / pct-encoded / sub-delims / ":" )
//     unreserved  = ALPHA / DIGIT / "-" / "." / "_" / "~"
//     sub-delims  = "!" / "$" / "&" / "'" / "(" / ")"
//                   / "*" / "+" / "," / ";" / "="
//
// It doesn't validate pct-encoded. The caller does that via func unescape.
func validUserinfo(s string) bool {
	for _, r := range s {
		if 'A' <= r && r <= 'Z' {
			continue
		}
		if 'a' <= r && r <= 'z' {
			continue
		}
		if '0' <= r && r <= '9' {
			continue
		}
		switch r {
		case '-', '.', '_', ':', '~', '!', '$', '&', '\'',
			'(', ')', '*', '+', ',', ';', '=', '%', '@':
			continue
		default:
			return false
		}
	}
	return true
}
