package main

import "fmt"

// BaseChannel 基础的proxy方法
type BaseChannel struct {
	liveFlag bool
	latency  int
}

func (channel *BaseChannel) isAlive() bool {
	return channel.liveFlag
}

func (channel *BaseChannel) getLatency() int {
	return channel.latency
}

func (channel *BaseChannel) setLatency(value int) {
	channel.latency = value
}

func proxyCheckTTL(channel IProxyChannel, url string) int {
	req := buildGetRequestFromURL(url)
	channel.request(req) // skip first one
	start := getTimestamp()
	channel.request(req) // skip first one
	ttl := int(getTimestamp()-start) / 1000000
	fmt.Printf("bench ttl for %s, and cost %d ms\n", url, ttl)
	channel.setLatency(ttl)
	return ttl
}

func findMinLatencyProxy(channels []IProxyChannel) IProxyChannel {
	if len(channels) == 0 {
		return nil
	}
	result := channels[0]
	for _, proxy := range channels {
		if result.getLatency() == 0 {
			result = proxy
		} else if result.getLatency() > proxy.getLatency() {
			result = proxy
		}
	}
	return result
}
