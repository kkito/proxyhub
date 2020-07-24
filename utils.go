package main

import (
	"log"
	"net"
	"net/http"

	"github.com/oschwald/geoip2-golang"
)

// 把来自proxy的request变成不同的request
func proxyRequest2Plain(req *http.Request) *http.Request {
	req.RequestURI = ""
	return req
}

func buildHTTPClient(httpTransport *http.Transport) *http.Client {
	return &http.Client{
		Transport: httpTransport,
		// 302 不会跳转，让客户端浏览器做跳转
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

}

func isCNHost(hostOrIP string) bool {
	countryCode := getCountryCodeByHostOrIP(hostOrIP)
	return countryCode == "CN"
}

func getCountryCodeByHostOrIP(hostOrIP string) string {
	addr := getIPFromHost(hostOrIP)
	if addr == nil {
		return ""
	}
	db, err := geoip2.Open("./data/GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer db.Close()
	// ip := net.ParseIP("81.2.69.142")
	record, err := db.Country(addr)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return record.Country.IsoCode
}

func getIPFromHost(host string) net.IP {
	addr, err := net.LookupIP(host)
	if err == nil && len(addr) > 0 {
		return addr[0]
	}
	// fmt.Println("Unknown host")
	return nil
}

func isInStringArray(arr []string, target string) bool {
	for _, x := range arr {
		if x == target {
			return true
		}
	}
	return false
}
