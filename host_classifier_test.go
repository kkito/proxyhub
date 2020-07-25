package main

import "testing"

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
