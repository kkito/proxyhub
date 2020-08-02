package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
)

func runLocalServer(req *http.Request) *http.Response {
	// TODO
	// download pem
	// list status
	if strings.Contains(req.URL.Host, "whwtaurl??") {
		fmt.Println("constainsss")
		// goproxy.ContentTypeText
		// http.StatusNotFound
		return goproxy.NewResponse(req,
			goproxy.ContentTypeHtml, http.StatusOK,
			"<h1>Don't waste your time!</h1>")

	}
	return nil
}
