package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8084", "proxy listen address")
	flag.Parse()
	setCustomCA()
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxyHub := buildProxyHubFromConfig()
	hostLRU := makeHostCheckLRU()
	proxyHub.execBenchmark()

	proxy.OnRequest().DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (retReq *http.Request, retRep *http.Response) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in f", r)
					// set return value
					retReq = req
					retRep = nil
				}
			}()
			resp := runLocalServer(req)
			if resp != nil {
				return req, resp
			}
			hcf := buildHostClassifierWithHostLRU(req.URL.Host, hostLRU)
			// if strings.Contains(req.URL.Host, "google") ||
			// 	strings.Contains(req.URL.Host, "youtube") ||
			// 	strings.Contains(req.URL.Host, "yt") ||
			// 	strings.Contains(req.URL.Host, "gsta") {
			if hcf.isWallBlock() {
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
	fmt.Println("start proxy by ", *addr)
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
