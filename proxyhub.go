package main

import (
	"encoding/json"
	"io/ioutil"
)

// ProxyHub 所有的proxy集中到一起
type ProxyHub struct {
	proxies []IProxyChannel
}

func (hub *ProxyHub) getProxies() []IProxyChannel {
	return hub.proxies
}

func (hub *ProxyHub) chooseChannel(hostDest *IHostDestClassifier) IProxyChannel {
	proxies := hub.getProxies()
	if len(proxies) > 0 {
		return proxies[0]
	}
	return nil
}

// ProxyConfig 使用情况具体数据
type ProxyConfig struct {
	Type    string `json:"type"`
	Address string `json:"address"`
	FQ      bool   `json:"fq"`
}

// ProxyConfigs 返回内容
type ProxyConfigs struct {
	Configs []ProxyConfig `json:"proxies"`
}

func buildProxyHubFromConfig() *ProxyHub {
	pcs := ProxyConfigs{}
	file, _ := ioutil.ReadFile("proxy.json")
	result := &pcs
	json.Unmarshal([]byte(file), result)

	proxyHub := ProxyHub{}
	for _, config := range result.Configs {
		if config.Type == "socks5" {
			channel := Socks5Channel{config.Address, nil}
			proxyHub.proxies = append(proxyHub.proxies, &channel)
		}
		// fmt.Println(config.Address)
	}
	return &proxyHub
}
