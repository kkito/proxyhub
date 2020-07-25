package main

import "strings"

// TODO 部分host可以LRU的方式缓存起来
// TODO 配置那些自定义网络，可配置化，比如公司的内部域名

// HostClassifier 域名分类器
type HostClassifier struct {
	host                    string
	internalHostsFromConfig []string
}

// 是否内部网络
func (hc *HostClassifier) isInternal() bool {
	interIps := []string{"127.", "10.", "192.168"}
	for _, check := range interIps {
		if strings.Contains(hc.host, check) {
			return true
		}
	}
	hc.initInternalHosts()
	if len(hc.internalHostsFromConfig) == 0 {
		return false
	}
	for _, check := range hc.internalHostsFromConfig {
		if strings.Contains(hc.host, check) {
			return true
		}
	}
	return false
}

// 是否国内网络
func (hc *HostClassifier) isCN() bool {
	return isCNHost(hc.host)
}

func (hc *HostClassifier) isWallBlock() bool {
	return !hc.isCN()
}

func (hc *HostClassifier) initInternalHosts() bool {
	if hc.internalHostsFromConfig == nil {
		config := getProxyHubConfig()
		if config.InternalHosts != nil {
			hc.internalHostsFromConfig = config.InternalHosts
		} else {
			empty := []string{}
			hc.internalHostsFromConfig = empty
		}
		return true
	}
	return false
}

func buildHostClassifier(host string) *HostClassifier {
	result := strings.Split(host, ":")
	hc := HostClassifier{host: result[0]}
	return &hc
}
