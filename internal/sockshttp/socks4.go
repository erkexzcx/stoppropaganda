package sockshttp

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

func SOCKS4(network, addr string, auth *Auth, forward Dialer) (*Socks4Proxier, error) {
	s := &Socks4Proxier{
		Host:   addr,
		Auth:   auth,
		Dialer: forward,
	}

	return s, nil
}

// Constants to choose which version of SOCKS protocol to use.
const (
	TypeSOCKS4 = iota
	TypeSOCKS4A
)

type Socks4Proxier struct {
	Proto   int
	Host    string
	Auth    *Auth
	Timeout time.Duration
	Dialer  Dialer
}

func (cfg *Socks4Proxier) Dial(network, forwardAddr string) (net.Conn, error) {
	socksType := cfg.Proto
	proxyAddr := cfg.Host

	// dial TCP
	conn, err := cfg.Dialer.Dial("tcp", proxyAddr)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			conn.Close()
		}
	}()

	// connection request
	host, port, err := splitHostPort16(forwardAddr)
	if err != nil {
		return nil, err
	}
	ip := net.IPv4(0, 0, 0, 1).To4()
	if socksType == TypeSOCKS4 {
		ip, err = lookupIP(host)
		if err != nil {
			return nil, err
		}
	}
	req := []byte{
		4,                          // version number
		1,                          // command CONNECT
		byte(port >> 8),            // higher byte of destination port
		byte(port),                 // lower byte of destination port (big endian)
		ip[0], ip[1], ip[2], ip[3], // special invalid IP address to indicate the host name is provided
		0, // user id is empty, anonymous proxy only
	}
	if socksType == TypeSOCKS4A {
		req = append(req, []byte(host+"\x00")...)
	}

	resp, err := cfg.sendReceive(conn, req)
	if err != nil {
		return nil, err
	} else if len(resp) != 8 {
		return nil, errors.New("server does not respond properly")
	}
	switch resp[1] {
	case 90:
		// request granted
	case 91:
		return nil, errors.New("socks4: 91 = rejected/failed")
	case 92:
		return nil, errors.New("socks4: 92 = client is not running identd (or not reachable from server)")
	case 93:
		return nil, errors.New("socks4: 93 = client's identd could not confirm the user ID in the request")
	default:
		return nil, fmt.Errorf("socks4: %d = ? unknown error", resp[1])
	}
	// clear the deadline before returning
	if err := conn.SetDeadline(time.Time{}); err != nil {
		return nil, err
	}
	return conn, nil
}

func (c *Socks4Proxier) sendReceive(conn net.Conn, req []byte) (resp []byte, err error) {
	if c.Timeout > 0 {
		if err := conn.SetWriteDeadline(time.Now().Add(c.Timeout)); err != nil {
			return nil, err
		}
	}
	_, err = conn.Write(req)
	if err != nil {
		return
	}
	resp, err = c.readAll(conn)
	return
}
func (c *Socks4Proxier) readAll(conn net.Conn) (resp []byte, err error) {
	resp = make([]byte, 512)
	if c.Timeout > 0 {
		if err := conn.SetReadDeadline(time.Now().Add(c.Timeout)); err != nil {
			return nil, err
		}
	}
	n, err := conn.Read(resp)
	resp = resp[:n]
	return
}

func splitHostPort16(addr string) (host string, port uint16, err error) {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return "", 0, err
	}
	portInt, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return "", 0, err
	}
	port = uint16(portInt)
	return
}

func lookupIP(host string) (net.IP, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}
	if len(ips) == 0 {
		return nil, fmt.Errorf("cannot resolve host: %s", host)
	}
	ip := ips[0].To4()
	if len(ip) != net.IPv4len {
		return nil, errors.New("ipv6 is not supported by SOCKS4")
	}
	return ip, nil
}
