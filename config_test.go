package main

import "testing"

func TestInternalHosts(t *testing.T) {
	config := getProxyHubConfigByConfigName("proxy_hub_test.json")
	if len(config.InternalHosts) != 3 {
		t.FailNow()
	}
	if !isInStringArray(config.InternalHosts, "kkito.cn") {
		t.FailNow()
	}
	if isInStringArray(config.InternalHosts, "whatelse") {
		t.FailNow()
	}
}
