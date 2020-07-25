package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
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

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8084", "proxy listen address")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxyHub := buildProxyHubFromConfig()
	hostLRU := makeHostCheckLRU()

	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		hcf := buildHostClassifierWithHostLRU(req.URL.Host, hostLRU)
		// if strings.Contains(req.URL.Host, "google") ||
		// 	strings.Contains(req.URL.Host, "youtube") ||
		// 	strings.Contains(req.URL.Host, "yt") ||
		// 	strings.Contains(req.URL.Host, "gsta") {
		if !hcf.isInternal() && !hcf.isCN() {
			println(req.URL.Host)
			channel := proxyHub.chooseChannel(hcf)
			if channel == nil {
				return req, nil
			}
			fmt.Println("NOT CN " + req.URL.Host)
			resp := channel.request(req)
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
