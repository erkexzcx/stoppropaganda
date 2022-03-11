// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockshttp

import (
	"bytes"
	"errors"
	"net"
	"strconv"
	"strings"
	"time"
)

func HTTP(network, addr string, forward Dialer) (*HttpProxier, error) {
	s := &HttpProxier{
		network: network,
		addr:    addr,
		forward: forward,
	}

	return s, nil
}

type HttpProxier struct {
	network, addr string
	forward       Dialer
	Timeout       time.Duration
}

func (s *HttpProxier) Dial(network, addr string) (net.Conn, error) {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	conn, err := s.forward.Dial(s.network, s.addr)
	if err != nil {
		return nil, err
	}
	closeConn := &conn
	defer func() {
		if closeConn != nil {
			(*closeConn).Close()
		}
	}()
	if s.Timeout > 0 {
		conn.SetDeadline(time.Now().Add(s.Timeout))
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.New("proxy_http: failed to parse port number: " + portStr)
	}
	if port < 1 || port > 0xffff {
		return nil, errors.New("proxy_http: port number out of range: " + portStr)
	}

	buf := make([]byte, 0, 6+len(host))
	buf = append(buf, "CONNECT "...)
	buf = append(buf, addr...)
	buf = append(buf, " HTTP 1.1\r\n"...)
	buf = append(buf, "User-agent: Mozilla/4.0\r\n\r\n"...)

	if _, err := conn.Write(buf); err != nil {
		return nil, errors.New("proxy_http: failed to write CONNECT to HTTP proxy at " + s.addr + ": " + err.Error())
	}
	if s.Timeout > 0 {
		conn.SetDeadline(time.Now().Add(s.Timeout))
	}
	buf = make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, errors.New("proxy_http: failed to read from HTTP proxy at " + s.addr + ": " + err.Error())
	}
	lines := bytes.Split(buf[:n], []byte("\n"))
	if len(lines) == 0 {
		return nil, errors.New("proxy_http: received fewer lines from HTTP proxy at " + s.addr)
	}
	response := string(lines[0])
	rcodes := strings.SplitN(response, " ", 3)
	if len(rcodes) < 3 {
		return nil, errors.New("proxy_http: received fewer header args from HTTP proxy at " + s.addr)
	}
	if !strings.HasPrefix(rcodes[0], "HTTP") {
		return nil, errors.New("proxy_http: it is not HTTP proxy at " + s.addr)
	}
	if !strings.HasPrefix(rcodes[1], "200") {
		return nil, errors.New("proxy_http: received error code " + rcodes[1] + " HTTP proxy at " + s.addr)
	}
	closeConn = nil
	return conn, nil
}
