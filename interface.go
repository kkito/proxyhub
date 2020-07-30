package main

import "net/http"

// IHostDestClassifier check host type
type IHostDestClassifier interface {
	// try using  "geoip:private", "geoip:cn" etc
	isInternal() bool
	isCN() bool
	isWallBlock() bool
}

// IProxyChannel a proxy channel can use
type IProxyChannel interface {
	// how to invoke?
	canFQ() bool
	request(r *http.Request) *http.Response
	getTTL() int             // ttl to check speed
	isAlive() bool           // check if proxy is alive
	checkTTL(url string) int // check ttl
}

// IProxyChannelBenchmark can choose a best proxy channel
type IProxyChannelBenchmark interface {
}

// procedure

// IProxyHub a hub to get channel
type IProxyHub interface {
	getProxies() []IProxyChannel
	chooseChannel(hostDest *IHostDestClassifier) *IProxyChannel
}

// IRequestUtil util for http.Request
type IRequestUtil interface {
	BuildClassifier() *IHostDestClassifier
}
