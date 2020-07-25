FROM golang:alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add git

# https://goproxy.io/
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY="https://goproxy.io,direct"
RUN go get github.com/oschwald/geoip2-golang
RUN go get github.com/elazarl/goproxy
