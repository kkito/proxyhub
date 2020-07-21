package main

import "net/http"

// 把来自proxy的request变成不同的request
func proxyRequest2Plain(req *http.Request) *http.Request {
	req.RequestURI = ""
	return req
}
