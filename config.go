package main

import (
	"encoding/json"
	"io/ioutil"
)

// ProxyConfig 使用情况具体数据
type ProxyConfig struct {
	Type    string `json:"type"`
	Address string `json:"address"`
	FQ      bool   `json:"fq"`
}

// ProxyHubConfig 返回内容
type ProxyHubConfig struct {
	Configs        []ProxyConfig `json:"proxies"`
	InternalHosts  []string      `json:"internal_hosts"`
	CheckTTLSite   string        `json:"check_ttl_site"`
	CheckTTLSiteFQ string        `json:"check_ttl_site_fq"`
}

func getProxyHubConfig() *ProxyHubConfig {
	return getProxyHubConfigByConfigName("proxy_hub.json")
}

func getProxyHubConfigByConfigName(configFile string) *ProxyHubConfig {
	pcs := ProxyHubConfig{}
	file, _ := ioutil.ReadFile(configFile)
	result := &pcs
	json.Unmarshal([]byte(file), result)
	return result
}
