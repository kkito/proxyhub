package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/elazarl/goproxy"
)

// visit http://proxy.hub/ to list pages
func runLocalServer(req *http.Request) *http.Response {
	// TODO
	// download pem
	// list status
	if strings.Contains(req.URL.Host, "proxy.hub") {
		fmt.Println("constainsss")
		// goproxy.ContentTypeText
		// http.StatusNotFound
		return goproxy.NewResponse(req,
			goproxy.ContentTypeHtml, http.StatusOK,
			pagePem())

	}
	return nil
}

func pagePem() string {
	pemContent := renderTpl("pem", "")
	return merge2Layout(pemContent)
}

func merge2Layout(content string) string {
	return renderTpl("layout", content)
}

func renderTpl(tplName string, content interface{}) string {
	t, err := template.ParseFiles("html/" + tplName + ".html")
	if err != nil {
		panic(err)
	}
	var result bytes.Buffer
	err = t.Execute(&result, content)
	if err != nil {
		panic(err)
	}
	return result.String()
}
