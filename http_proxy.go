package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// HTTPChannel 实现 http proxy
type HTTPChannel struct {
	BaseChannel
}

// GetType get type
func (*HTTPChannel) GetType() string {
	return "HTTP"
}

func (*HTTPChannel) getTTL() int {
	return 0
}

func (*HTTPChannel) isAlive() bool {
	return true
}

func (channel *HTTPChannel) checkTTL(url string) int {
	return proxyCheckTTL(channel, url)
}

func (*HTTPChannel) canFQ() bool {
	return true
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
