package main

import "testing"

func TestFindMinLatency(t *testing.T) {
	ch1 := &HTTPChannel{}
	ch1.address = "test1"
	ch1.setLatency(0)
	ch2 := &HTTPChannel{}
	ch2.address = "test1"
	ch2.setLatency(0)
	items := []IProxyChannel{ch1, ch2}
	result := findMinLatencyProxy(items)
	if result != ch2 {
		t.Fail()
	}

	ch2.setLatency(10)
	result = findMinLatencyProxy(items)
	if result != ch2 {
		t.Fail()
	}

	ch1.setLatency(5)
	result = findMinLatencyProxy(items)
	if result != ch1 {
		t.Fail()
	}
}
