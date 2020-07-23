package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// HTTPChannel 实现 http proxy
type HTTPChannel struct {
	address string // eg "127.0.0.1:1887"
}

func (*HTTPChannel) canFQ() bool {
	return false
}

func (channel *HTTPChannel) request(req *http.Request) *http.Response {
	req = proxyRequest2Plain(req)
	proxyURL, err := url.Parse("http://" + channel.address)
	// fmt.Println("from " + channel.address)
	if err != nil {
		return nil
	}
	httpTransport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	httpClient := buildHTTPClient(httpTransport)
	// set our socks5 as the dialer
	resp, err := httpClient.Do(req)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		return nil
	}
	return resp

}
