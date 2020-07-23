# proxyhub

a hub for lots of proxy 

depends on 

* `https://github.com/elazarl/goproxy`
* `github.com/oschwald/geoip2-golang`

### start server

`go run .`

`go test -v ./...`

### features

* proxy for http & https
* support multi proxies
* choose a best proxy to use
  * period benchmark
* support different kind of proxy
  * socks5
  * http proxy
  * others maybe ss or v2rays
* proxy has a attr it can FQ or fast etc
* check host or ip to use special proxy 
    * internal without proxy
    * FQ proxy
    * domestic fast way

#### lower prority

* download ca certification from a special url to install 
* a web interface to check all access
  * config a basic auth
  * basic auth username and password can config
* can pass a ca cert to start
* proxy can set a weight
* some host can set to special strategy

#### geoip2

download data from https://www.maxmind.com/en/accounts/364812/geoip/downloads

only use GeoLite2 Country to check if CN