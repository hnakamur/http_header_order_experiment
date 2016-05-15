package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "127.0.0.1:8008", "listen address")
	var upstream string
	flag.StringVar(&upstream, "upstream", "http://127.0.0.1:8080", "upstream URL")
	flag.Parse()

	upstreamURL, err := url.Parse(upstream)
	if err != nil {
		log.Fatal(err)
	}
	handler := httputil.NewSingleHostReverseProxy(upstreamURL)
	http.Handle("/", handler)
	http.ListenAndServe(addr, nil)
}
