package main

import (
	"fmt"
	"testing"
)

func TestIfInitWithNil(t *testing.T) {
	hc := HostClassifier{host: "test.com"}
	// init struct the internal hosts is  nil
	if hc.internalHostsFromConfig != nil {
		t.FailNow()
	}
}

func TestIsInternal(t *testing.T) {
	hc := HostClassifier{
		host:                    "test.com",
		internalHostsFromConfig: []string{"test"},
	}
	if !hc.isInternal() {
		t.Fail()
	}
	hc.host = "other.com"
	if hc.isInternal() {
		t.Fail()
	}

}

func TestArrayInit(t *testing.T) {
	fort := make([]string, 3)
	if fort[0] != "" {
		t.FailNow()
	}
}

func TestHostCheckLRUInit(t *testing.T) {
	hcl := makeHostCheckLRU()
	result := hcl.hasHost("test.com")
	if result {
		t.Fail()
	}
	hcl.pushHost("test.com", false)
	result = hcl.hasHost("test.com")
	if !result {
		t.Fail()
	}
	result = hcl.isMeet("test.com")
	if result {
		t.Fail()
	}

}

func TestGetMinClearTimestamp(t *testing.T) {
	hcl := makeHostCheckLRU()
	hcl.maxSize = 4
	v := hcl.pushHost("test.com", true)
	v.lastCheckTimeStamp = int64(1)

	v = hcl.pushHost("test1.com", false)
	v.lastCheckTimeStamp = int64(4)
	v = hcl.pushHost("test3.com", false)
	v.lastCheckTimeStamp = int64(3)
	v = hcl.pushHost("test4.com", false)
	v.lastCheckTimeStamp = int64(2)

	hcl.removeSize = 3
	result := hcl.getMinClearTimestamp()
	if result != 3 {
		t.Fail()
	}

	hcl.removeSize = 1
	result = hcl.getMinClearTimestamp()
	if result != 1 {
		t.Fail()
	}

	hcl.removeSize = 3
	hcl.removeOldests()
	t.Logf(fmt.Sprintf("%d size", len(hcl.hostTimeMap)))
	if len(hcl.hostTimeMap) != 1 {
		t.Fail()
	}

}
