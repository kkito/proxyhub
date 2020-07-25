  FROM golang:alpine
  RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
  RUN apk add git

#   RUN export CGO_ENABLED=0
  RUN export GO111MODULE=on
  RUN export GOPROXY=https://goproxy.io
  RUN go get github.com/oschwald/geoip2-golang
  RUN go get github.com/elazarl/goproxy
