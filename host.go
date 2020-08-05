package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	"github.com/elazarl/goproxy"
)

// 标准的返回page方法返回内容
type pageResp struct {
	contentType  string
	responseCode int
	content      string
	filename     string
}

// map of path to pageFunc
var pathMap = map[string]func(*http.Request) pageResp{
	"/test":         pageTest,
	"/pem":          pagePem,
	"/pem/download": pagePemDownload,
}

var contentTypeFile = "application/octet-stream"

// visit http://proxy.hub/ to list pages
func runLocalServer(req *http.Request) *http.Response {
	// TODO
	// download pem
	// list status

	if strings.Contains(req.URL.Host, "proxy.hub") {
		pageFunc, ok := pathMap[req.URL.Path]
		if !ok {
			pageFunc = pageNotFound
		}
		ret := pageFunc(req)
		resp := goproxy.NewResponse(req,
			ret.contentType, ret.responseCode,
			ret.content)
		if ret.contentType == contentTypeFile {
			resp.Header.Add("Content-Disposition",
				fmt.Sprintf("attachment;filename=%s", ret.filename))
			buf := bytes.NewBufferString(ret.content)
			resp.ContentLength = int64(buf.Len())
			resp.Body = ioutil.NopCloser(buf)

		}
		return resp

	}
	return nil
}

func pageTest(req *http.Request) pageResp {
	return _htmlReturn("<h1> test page</h1>")
}

func pageNotFound(req *http.Request) pageResp {
	return _htmlReturn("<h1> NOT FOUND</h1>")
}

func pagePem(req *http.Request) pageResp {
	pemContent := renderTpl("pem", "")
	return _htmlReturn(pemContent)
}

func pagePemDownload(*http.Request) pageResp {
	return pageResp{
		contentTypeFile,
		http.StatusOK,
		string(caCert),
		"proxyhub.pem",
	}
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

func _htmlReturn(html string) pageResp {
	return pageResp{
		goproxy.ContentTypeHtml,
		http.StatusOK,
		merge2Layout(html),
		"",
	}
}
