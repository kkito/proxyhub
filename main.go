package main

import (
	"flag"
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
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if req.URL.Scheme == "https" {
			// req.URL.Scheme = "http"
			println(req.URL.Scheme)
			println(req.URL.Host)
		}
		return req, nil
	})
	proxy.Verbose = *verbose
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
