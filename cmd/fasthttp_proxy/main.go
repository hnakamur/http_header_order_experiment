package main

import (
	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

// This example was copied from https://github.com/valyala/fasthttp/issues/64#issuecomment-194880575
// and modified.

type reverseProxy struct {
	proxyClient *fasthttp.HostClient
}

func newReverseProxy(hostAddr string) *reverseProxy {
	return &reverseProxy{
		proxyClient: &fasthttp.HostClient{
			Addr: hostAddr,
		},
	}
}

func (p *reverseProxy) Handle(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	resp := &ctx.Response
	p.prepareRequest(req)
	if err := p.proxyClient.Do(req, resp); err != nil {
		ctx.Logger().Printf("error when proxying the request: %s", err)
	}
	p.postprocessResponse(resp)
}

func (p *reverseProxy) prepareRequest(req *fasthttp.Request) {
	// do not proxy "Connection" header.
	//req.Header.Del("Connection")
	// strip other unneeded headers.

	// alter other request params before sending them to upstream host
}

func (p *reverseProxy) postprocessResponse(resp *fasthttp.Response) {
	resp.Header.Add("Via", "my fasthttp proxy")

	// strip other unneeded headers

	// alter other response data if needed
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "127.0.0.1:8008", "listen address")
	var upstreamAddr string
	flag.StringVar(&upstreamAddr, "upstream-addr", "127.0.0.1:8080", "upstream address")
	flag.Parse()

	rp := newReverseProxy(upstreamAddr)

	// Start the server listening for incoming requests on the given address.
	//
	// ListenAndServe returns only on error, so usually it blocks forever.
	if err := fasthttp.ListenAndServe(addr, rp.Handle); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
