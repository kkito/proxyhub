package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/proxy"
)

// Socks5Channel 实现socks5 proxy
type Socks5Channel struct {
	address string // eg "127.0.0.1:1887"
	dialer  *proxy.Dialer
}

func (channel *Socks5Channel) getDialer() *proxy.Dialer {
	if channel.dialer == nil {
		dialer, err := proxy.SOCKS5("tcp", channel.address, nil, proxy.Direct)
		// dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1887", nil, proxy.Direct)
		if err != nil {
			fmt.Println(err)
		} else {
			channel.dialer = &dialer
		}
	}
	return channel.dialer
}

func (*Socks5Channel) canFQ() bool {
	return false
}

func (channel *Socks5Channel) request(req *http.Request) *http.Response {
	dialer := channel.getDialer()
	if dialer == nil {
		return nil
	}
	req = proxyRequest2Plain(req)
	httpTransport := &http.Transport{}
	httpClient := buildHTTPClient(httpTransport)
	// set our socks5 as the dialer
	httpTransport.Dial = (*dialer).Dial
	resp, err := httpClient.Do(req)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		return nil
	}
	return resp

}
