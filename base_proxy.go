package main

import "fmt"

// BaseChannel 基础的proxy方法
type BaseChannel struct {
	liveFlag bool
	latency  int
	address  string // eg "127.0.0.1:1887"
}

// GetAddress method
func (channel *BaseChannel) GetAddress() string {
	return channel.address
}

func (channel *BaseChannel) isAlive() bool {
	return channel.liveFlag
}

// GetLatency in millseconds
func (channel *BaseChannel) GetLatency() int {
	return channel.latency
}

func (channel *BaseChannel) setLatency(value int) {
	channel.latency = value
}

func (channel *BaseChannel) setLiveFlag(value bool) {
	channel.liveFlag = value
}

func proxyCheckTTL(channel IProxyChannel, url string) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			channel.setLiveFlag(false)
			ret = -1
			return
		}
	}()
	channel.setLiveFlag(true)
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
		if !proxy.isAlive() {
			continue
		}
		if result.GetLatency() == 0 {
			result = proxy
		} else if result.GetLatency() > proxy.GetLatency() {
			result = proxy
		}
	}
	return result
}
