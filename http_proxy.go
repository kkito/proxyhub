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

func (*HTTPChannel) getTTL() int {
	return 0
}

func (*HTTPChannel) isAlive() bool {
	return true
}

func (channel *HTTPChannel) checkTTL(url string) int {
	req := buildGetRequestFromURL(url)
	channel.request(req) // skip first one
	start := getTimestamp()
	channel.request(req) // skip first one
	ttl := int(getTimestamp()-start) / 1000000
	fmt.Println(start)
	fmt.Printf("bench ttl for %s, and cost %d\n", url, ttl)
	return ttl
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
