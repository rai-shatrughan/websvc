package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"websvc/gateway/handler"
)

var pMap = make(map[string]string)
var scheme = "http"

func init() {
	pMap["127.0.0.1:8001"] = "/api/v1/device/"
	pMap["127.0.0.1:8002"] = "/api/v1/timeseries/"
}

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	sema := make(chan struct{}, 2000)
	return func(w http.ResponseWriter, r *http.Request) {
		sema <- struct{}{}
		defer func() { <-sema }()
		proxy.Transport = &handler.Transport{}
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	for k, v := range pMap {
		// initialize a reverse proxy and pass the actual backend server url here
		proxy, err := NewProxy(scheme + "://" + k)
		if err != nil {
			panic(err)
		}

		// handle all requests to your server using the proxy
		go http.HandleFunc(v, ProxyRequestHandler(proxy))
	}
	log.Printf("Starting Gateway at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
