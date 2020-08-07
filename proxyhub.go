package main

import "fmt"

// ProxyHub 所有的proxy集中到一起
type ProxyHub struct {
	proxies         []IProxyChannel
	config          *ProxyHubConfig
	UseCount        int // the amount of use proxy
	ReBenchmarkStep int // after step to rebenchmark all proxies
}

func (hub *ProxyHub) getProxies() []IProxyChannel {
	return hub.proxies
}

func (hub *ProxyHub) chooseChannel(hostDest IHostDestClassifier) IProxyChannel {
	proxies := hub.getProxies()
	if hostDest != nil && hostDest.isWallBlock() {
		proxies = hub.getAllCanFQChannels()
	}
	hub.UseCount++
	if hub.UseCount%hub.ReBenchmarkStep == 0 {
		// TODO run it in deamon
		hub.execBenchmark()
	}
	return findMinLatencyProxy(proxies)
}

func (hub *ProxyHub) getAllCanFQChannels() (ret []IProxyChannel) {
	if hub.proxies == nil {
		return
	}
	for _, proxy := range hub.proxies {
		if proxy.canFQ() {
			ret = append(ret, proxy)
		}
	}
	return
}

// execBenchmark for a long period or something error happened
func (hub *ProxyHub) execBenchmark() {
	fmt.Println("start run benchmark!!")
	for _, proxy := range hub.proxies {
		site := hub.config.CheckTTLSite
		fmt.Println(proxy.canFQ())
		if proxy.canFQ() {
			site = hub.config.CheckTTLSiteFQ
			fmt.Println(site)
		}
		proxy.checkTTL(site)
	}
}

func buildProxyHubFromConfig() *ProxyHub {
	result := getProxyHubConfig()

	proxyHub := ProxyHub{
		config:          result,
		UseCount:        0,
		ReBenchmarkStep: 2000,
	}

	for _, config := range result.Configs {
		if config.Type == "socks5" {
			channel := Socks5Channel{dialer: nil, alive: false, ttl: 0}
			channel.address = config.Address
			proxyHub.proxies = append(proxyHub.proxies, &channel)
		}
		if config.Type == "http" {
			channel := HTTPChannel{}
			channel.address = config.Address
			proxyHub.proxies = append(proxyHub.proxies, &channel)
		}
		// fmt.Println(config.Address)
	}
	return &proxyHub
}
