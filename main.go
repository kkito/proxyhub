package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/elazarl/goproxy"
	"golang.org/x/net/proxy"
)

func doProxy(req *http.Request, ctx *goproxy.ProxyCtx, requestUtil *IRequestUtil, proxyHub *IProxyHub) (*http.Request, *http.Response) {
	classifilter := (*requestUtil).BuildClassifier()
	if (*classifilter).isInternal() {
		return req, nil
	}

	// if not internal do something else
	proxyChanel := (*proxyHub).chooseChannel(classifilter)
	res := (*proxyChanel).request(req)
	return req, res
}

func socksProxy(req *http.Request) *http.Response {
	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1887", nil, proxy.Direct)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		return nil
	}
	proxyGet := req.Header.Get("GET")
	println(proxyGet)
	// reader, err := req.GetBody()
	// if err != nil {
	// 	// panic(err)
	// 	fmt.Println(err)
	// 	return nil
	// }
	// fmt.Println(reader)
	fmt.Println("======")
	fmt.Println(req.Host)
	fmt.Println(req.URL)
	reqNew, err := http.NewRequest(req.Method, req.URL.String(), nil)
	reqNew.Header = req.Header
	reqNew.Body = req.Body
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create request:", err)
		os.Exit(2)
	}
	req = reqNew

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

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8084", "proxy listen address")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if strings.Contains(req.URL.Host, "google") ||
			strings.Contains(req.URL.Host, "youtube") ||
			strings.Contains(req.URL.Host, "yt") ||
			strings.Contains(req.URL.Host, "gsta") {
			println(req.URL.Host)
			resp := socksProxy(req)
			return req, resp
		}
		// if req.URL.Scheme == "https" {
		// 	// req.URL.Scheme = "http"
		// 	println(req.URL.Scheme)
		// }
		return req, nil
	})
	proxy.Verbose = *verbose
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
