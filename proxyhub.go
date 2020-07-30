package main

// ProxyHub 所有的proxy集中到一起
type ProxyHub struct {
	proxies []IProxyChannel
	config  *ProxyHubConfig
}

func (hub *ProxyHub) getProxies() []IProxyChannel {
	return hub.proxies
}

func (hub *ProxyHub) chooseChannel(hostDest IHostDestClassifier) IProxyChannel {
	proxies := hub.getProxies()
	if hostDest != nil && hostDest.isWallBlock() {
		proxies = hub.getAllCanFQChannels()
	}
	if len(proxies) > 0 {
		// fmt.Println("proxy found!")
		return proxies[0]
	}
	return nil
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
	for _, proxy := range hub.proxies {
		site := hub.config.CheckTTLSite
		if proxy.canFQ() {
			site = hub.config.CheckTTLSiteFQ
		}
		proxy.checkTTL(site)
	}
}

func buildProxyHubFromConfig() *ProxyHub {
	result := getProxyHubConfig()

	proxyHub := ProxyHub{config: result}
	for _, config := range result.Configs {
		if config.Type == "socks5" {
			channel := Socks5Channel{config.Address, nil}
			proxyHub.proxies = append(proxyHub.proxies, &channel)
		}
		if config.Type == "http" {
			channel := HTTPChannel{config.Address}
			proxyHub.proxies = append(proxyHub.proxies, &channel)
		}
		// fmt.Println(config.Address)
	}
	return &proxyHub
}
