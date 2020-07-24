package main

import "strings"

// TODO 部分host可以LRU的方式缓存起来

// HostClassifier 域名分类器
type HostClassifier struct {
	host string
}

// 是否内部网络
func (hc *HostClassifier) isInternal() bool {
	return false
}

// 是否国内网络
func (hc *HostClassifier) isCN() bool {
	return isCNHost(hc.host)
}

func (hc *HostClassifier) isWallBlock() bool {
	return !hc.isCN()
}

func buildHostClassifier(host string) *HostClassifier {
	result := strings.Split(host, ":")
	hc := HostClassifier{result[0]}
	return &hc
}
