package main

import "testing"

func TestFromJson(t *testing.T) {
	result := buildProxyHubFromConfig()
	if len(result.getProxies()) != 1 {
		t.FailNow()
	}
	proxy := result.chooseChannel(nil)
	if proxy.canFQ() {
		t.FailNow()
	}
}
