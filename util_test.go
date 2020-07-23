package main

import "testing"

func TestIsCNHost(t *testing.T) {
	isCNHost("www.baidu.com")
}

func TestGetIPFromHost(t *testing.T) {
	result := getIPFromHost("www.baidu.com")
	t.Logf(result.String())
	t.Logf(getCountryCodeByHostOrIP("www.baidu.com"))
	result = getIPFromHost("114.114.114.114")
	t.Logf(result.String())
	t.Logf(getCountryCodeByHostOrIP("114.114.114.114"))
	t.Logf(getCountryCodeByHostOrIP("kkito.cn"))
	if isCNHost("kkito.cn") {
		t.FailNow()
	}
}
