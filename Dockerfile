  FROM golang:alpine

  RUN go get github.com/oschwald/geoip2-golang
  RUN go get github.com/elazarl/goproxy
