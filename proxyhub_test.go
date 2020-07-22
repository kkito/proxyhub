package main

import "testing"

func TestFromJson(t *testing.T) {
	result := buildProxyHubFromConfig()
	// t.Log(len(result.getProxies()))
	if len(result.getProxies()) <= 0 {
		t.FailNow()
	}
	proxy := result.chooseChannel(nil)
	if proxy.canFQ() {
		t.FailNow()
	}
}
