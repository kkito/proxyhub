package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

// TODO 部分host可以LRU的方式缓存起来
// TODO 配置那些自定义网络，可配置化，比如公司的内部域名

var lock = sync.Mutex{}

// HostClassifier 域名分类器
type HostClassifier struct {
	host                    string
	internalHostsFromConfig []string
	hostLRU                 *HostCheckLRU
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
	if hc.hostLRU != nil {
		if hc.hostLRU.hasHost(hc.host) {
			// fmt.Println("cache host for " + hc.host)
			hc.hostLRU.updateHost(hc.host)
			return hc.hostLRU.isMeet(hc.host)
		}
	}
	result := isCNHost(hc.host)
	if hc.hostLRU != nil {
		hc.hostLRU.pushHost(hc.host, result)
	}
	return result
}

func (hc *HostClassifier) isWallBlock() bool {
	return !hc.isCN() && !hc.isInternal()
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

func buildHostClassifierWithHostLRU(host string, hostLRU *HostCheckLRU) *HostClassifier {
	result := buildHostClassifier(host)
	result.hostLRU = hostLRU
	return result
}

// ================ host check LRU ==========

// HostCheckValue the value
type HostCheckValue struct {
	value              bool
	lastCheckTimeStamp int64
}

// IsTrue if hit
func (hcv *HostCheckValue) IsTrue() bool {
	return hcv.value
}

// HostCheckLRU keep exist
type HostCheckLRU struct {
	hostTimeMap map[string]*HostCheckValue
	maxSize     int
	removeSize  int
}

func makeHostCheckLRU() *HostCheckLRU {
	result := HostCheckLRU{
		hostTimeMap: make(map[string]*HostCheckValue),
		maxSize:     2048,
		removeSize:  256,
	}
	return &result
}

func (hcl *HostCheckLRU) hasHost(host string) bool {
	_, ok := hcl.hostTimeMap[host]
	return ok
}

func (hcl *HostCheckLRU) isMeet(host string) bool {
	value := hcl.hostTimeMap[host]
	return value.value
}

func (hcl *HostCheckLRU) pushHost(host string, value bool) *HostCheckValue {
	if hcl.isFull() {
		hcl.removeOldests()
	}
	result := HostCheckValue{value, getTimestamp()}
	hcl.hostTimeMap[host] = &result
	return &result
}

func (hcl *HostCheckLRU) updateHost(host string) (ret bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in updateHost", r)
			// return false if recover
			ret = false
		}
	}()
	lock.Lock()
	value, ok := hcl.hostTimeMap[host]
	result := false
	if ok {
		hcl.hostTimeMap[host] = &HostCheckValue{value.value, getTimestamp()}
		result = true
	}
	lock.Unlock()
	return result
}

func (hcl *HostCheckLRU) isFull() bool {
	return len(hcl.hostTimeMap) >= hcl.maxSize
}

func (hcl *HostCheckLRU) removeOldests() {

	minTimestamp := hcl.getMinClearTimestamp()
	for host, v := range hcl.hostTimeMap {
		if v.lastCheckTimeStamp <= minTimestamp {
			delete(hcl.hostTimeMap, host)
		}
	}
}
func (hcl *HostCheckLRU) getMinClearTimestamp() int64 {
	tss := make([]int64, hcl.maxSize)
	for _, v := range hcl.hostTimeMap {
		tss = append(tss, v.lastCheckTimeStamp)
	}
	sort.Slice(tss, func(i, j int) bool { return tss[i] > tss[j] })
	return tss[hcl.maxSize-hcl.removeSize]
}
