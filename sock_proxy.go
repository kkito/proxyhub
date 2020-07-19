package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/proxy"
)

func getDialer() proxy.Dialer {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1887", nil, proxy.Direct)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		panic(err)
	}
	return dialer

}

func copyRequest(req *http.Request) *http.Request {
	// https://golang.org/pkg/net/http/#Request
	fmt.Println("sssss")
	fmt.Println(req.Host)
	fmt.Println(req.RequestURI)
	req.RequestURI = ""
	results := strings.Split(req.RequestURI, "://")
	if len(results) <= 1 {
		return req
	}
	parts := strings.Split(results[1], "/")
	fmt.Println(parts)
	return req

}

func socksProxy(req *http.Request, dialer proxy.Dialer) *http.Response {
	if dialer == nil {
		dialer = getDialer()
	}
	// create a socks5 dialer
	proxyGet := req.Header.Get("GET")
	println(proxyGet)
	// reader, err := req.GetBody()
	// if err != nil {
	// 	// panic(err)
	// 	fmt.Println(err)
	// 	return nil
	// }
	// fmt.Println(reader)
	// fmt.Println("======")
	// fmt.Println(req.Host)
	// fmt.Println(req.URL)
	// reqNew, err := http.NewRequest(req.Method, req.URL.String(), nil)
	// reqNew.Header = req.Header
	// reqNew.Body = req.Body
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "can't create request:", err)
	// 	os.Exit(2)
	// }
	// req = reqNew

	req = copyRequest(req)

	fmt.Println("======")
	fmt.Println(req.Host)
	fmt.Println(req.URL)
	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	resp, err := httpClient.Do(req)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		return nil
	}
	println("return valid resp")
	return resp
}
