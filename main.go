package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

var caKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA1tpwKN2o3eOPhpkVFQj++pJpe7SYeoeGjkxrPdPV4tqAip10
xcwt+SQMkgWmTo2kf6UZjJ0JmfXUBkgIATHC7doUnvkSEQAUpsc+xr3pXSyodFFJ
5EPaaUbmrRG6fMfEP3RolSmJaQ9L11BDly1zr7hSUiAaCXuyuXh0It8uxf2gyxEg
b9YBwfbOBKMC/JUrQJjGoNOMyUJGbmY99bLAuOpENI8rEOoDfA0cULC2qVeUO2yP
/N0Zi/2AJsdFItnBzA0cmi8gAa7FVfd0+ktj/Cgw2OCQEI5kk/RZrJ3AXIJQEnB1
daIRp5sd8h4mg5R/HjwPauMqoMLQ/qGLivsTyQIDAQABAoIBACnDThfzdjajXCu6
p+lt1TpZqV1dbmIq49HXMuVSvvmYpXHMR32HQcxy4Gql5HzSdY5GRmAZylr9+Ne7
uqqYxJ49TNLV8VgSnvEIO8Cf+7Ob0abCgk13jwX1vTMZBhnpLtFyzD1qVIZybbiN
poXFVP5sIrxJ9yWuwL/ilRiOL2ZWmm7oFuPuv8WCVtdIT25yT7x+PnbWIO+34U5t
gtpMBlutAYLDZ4LXRMSu9w81HfcC0z6KZLFiT46CAijQLVsDpqQ+3cKhFQutCFiH
bvO8q4ku6Pq5WXgP8JrnEyYuiEvNmoZkizEiR5ZNR/bc930yKQBLiXxt7U8QjuOB
X4cVqY0CgYEA7+gfBlf+2mfB+1uAebn4xeTWlWFgJJMqGE0ylCdx7w/QfmeQ3xbd
3DF1hgT+h7N2qH7aYSGQJkhnXEgZjlFu4v9NuoN6uez7jvjstcIJ+0+V3Jvf/wlu
wKz//83NIo6wHZymJ28Usi2ycGN5exi8LkWaLz7mA5isg6wZ80vetVMCgYEA5UQU
AfAEk3uF3LIx726V6+Io0oPI292Thwwa3AdjT20G8G2zbz1ptQcdq9W69jo/pazn
UvkkifIbXJVsN/g+Cyc6h6hm9MkokjriNEELga0a79O75TiH513oZJ3PviGfWSUA
/mR9kC054S4qMut8LDj7EKnhbPHncIyKd74ScvMCgYEApIX7QM536C/fyBEoOlJf
WNdmkWsGFA8YfzHxzch/SgL+aLF9mICGBculXRNvuoBIj6Tu/k2WHarpt096ty8B
bIrJM2+XaooquhHbw4oebkpV68S6CJzfZyM9LKBmXZydCrGzALgc6VSNWqXdWZ3M
766r3lq0QyMgq+wKn98YDCECgYBv6b52f6pML9TVOWrkvK+UOI7CzC+lG9Ei8AbA
dx7EK61iebpIR8ss4e9a7PbZsO9WuUMmHpX2fGdc11e/Ln9ixGBuzgaL4RHb58B5
z3KFd3GZtlqW9vRoPU/upZY98n2tb0G/7F/anCkPwZA50PeJQrtTlAmFO8RFDWWe
M/sffwKBgEW5AY2cTJ1HC9gx38oCY23ZAlAlUae8MiSmAwYJFGupXxk1Y6ifp3Rq
S1FGcSGGVRboLPOfI33lXj/YwJch2cqIDgSzEwOjcQ38GhktlJvMpCo5Ro+BD9bF
XaGZBCGiq4l+QESElj8lxGbzwcH9zAh0lP+B4xOxGfBVmyJrelGR
-----END RSA PRIVATE KEY-----`)
var caCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDiDCCAnACCQDY6/vQFf91DDANBgkqhkiG9w0BAQsFADCBhTELMAkGA1UEBhMC
Q04xCzAJBgNVBAgMAlNIMREwDwYDVQQHDAhzaGFuZ2hhaTERMA8GA1UECgwIa2tp
dG8uY24xDjAMBgNVBAsMBWtraXRvMQ4wDAYDVQQDDAVra2l0bzEjMCEGCSqGSIb3
DQEJARYUa2tpdG9ra2l0b0BnbWFpbC5jb20wHhcNMjAwODAyMDIzMjAyWhcNMzAw
NzMxMDIzMjAyWjCBhTELMAkGA1UEBhMCQ04xCzAJBgNVBAgMAlNIMREwDwYDVQQH
DAhzaGFuZ2hhaTERMA8GA1UECgwIa2tpdG8uY24xDjAMBgNVBAsMBWtraXRvMQ4w
DAYDVQQDDAVra2l0bzEjMCEGCSqGSIb3DQEJARYUa2tpdG9ra2l0b0BnbWFpbC5j
b20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDW2nAo3ajd44+GmRUV
CP76kml7tJh6h4aOTGs909Xi2oCKnXTFzC35JAySBaZOjaR/pRmMnQmZ9dQGSAgB
McLt2hSe+RIRABSmxz7GveldLKh0UUnkQ9ppRuatEbp8x8Q/dGiVKYlpD0vXUEOX
LXOvuFJSIBoJe7K5eHQi3y7F/aDLESBv1gHB9s4EowL8lStAmMag04zJQkZuZj31
ssC46kQ0jysQ6gN8DRxQsLapV5Q7bI/83RmL/YAmx0Ui2cHMDRyaLyABrsVV93T6
S2P8KDDY4JAQjmST9FmsncBcglAScHV1ohGnmx3yHiaDlH8ePA9q4yqgwtD+oYuK
+xPJAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAL94MTGslBPcTnrVMsxmnDD8bE+i
eRV/W4EOuxVGgkH60iej/tqEEdWDStsj7qRos/o0JECE6pWJfS/x9tlMvS8bKQ2I
LuytO+kDd7lmfRpkbb3p1DCSywIex2KU3mT9l+VEn2+CLSNG74wr+/d5ID8x5RcP
l/CC0k/RkEarisNSQp31Yh+zff9ENVe+DdrCedU414aQR8nHSiY0jAG2krWBPoK8
3WJ2KOqR7q4Ri2cKEm8Tx/5e6knXkkHG2ge68WxQDhe5bLuWATLX/myu15GsxFPv
wO9xUkPNtd1WwinjjkwreEMm7bfkMH8cTlaQ0LLgUbXQvpU616uTd5jnHMU=
-----END CERTIFICATE-----`)

func setCustomCA() {
	// 生成 ca
	// openssl genrsa -out cakey.pem 2048
	// openssl req -new -x509 -key cakey.pem -out cacert.pem -days 3650

	goproxyCa, err := tls.X509KeyPair(caCert, caKey)
	if err != nil {
		panic(err)
	}
	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
		panic(err)
	}
	goproxy.GoproxyCa = goproxyCa
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
}

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8084", "proxy listen address")
	flag.Parse()
	// setCustomCA()
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxyHub := buildProxyHubFromConfig()
	hostLRU := makeHostCheckLRU()
	proxyHub.execBenchmark()

	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (retReq *http.Request, retRep *http.Response) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
				// set return value
				retReq = req
				retRep = nil
			}
		}()
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
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
